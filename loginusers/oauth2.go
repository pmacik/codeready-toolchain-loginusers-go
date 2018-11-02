package loginusers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	common "github.com/pmacik/loginusers-go/common"

	"github.com/google/uuid"
	"github.com/tebeka/selenium"
)

// LoginUsersOAuth2 attempts to login into CodeReady Toolchain (former Openshift.io)
func LoginUsersOAuth2(authServerAddress string, userName string, userPassword string) (*Tokens, error) {
	wd, service := initSelenium()

	defer service.Stop()
	defer wd.Quit()

	clientID := common.Getenv("AUTH_CLIENT_ID", "740650a2-9c44-4db5-b067-a3d1b2cd2d01")
	redirectURL := fmt.Sprintf("%s/api/status", authServerAddress)
	state, _ := uuid.NewUUID()

	startURL := fmt.Sprintf("%s/api/authorize?response_type=code&client_id=%s&redirect_uri=%s&state=%s", authServerAddress, clientID, redirectURL, state.String())

	log.Printf("open-login-page...")
	if err := wd.Get(startURL); err != nil {
		return nil, fmt.Errorf("failed to open URL: '%s'", startURL)
	}

	findElementBy(wd, selenium.ByID, "kc-login")

	log.Printf("get-code...")
	sendKeysToElementBy(wd, selenium.ByID, "username", userName)
	elem := findElementBy(wd, selenium.ByID, "password")
	sendKeysToElement(elem, userPassword)
	submitElement(elem)

	wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		currentURL, _ := wd.CurrentURL()
		return strings.Contains(currentURL, state.String()), nil
	}, 10000)

	currentURL, _ := wd.CurrentURL()
	u, err := url.Parse(currentURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse URL: '%s'", currentURL)
	}
	code := u.Query()["code"]
	log.Printf("get-token...")
	resp, err := http.PostForm(
		fmt.Sprintf("%s/api/token", authServerAddress),
		url.Values{
			"grant_type":   {"authorization_code"},
			"client_id":    {clientID},
			"code":         code,
			"redirect_uri": {redirectURL},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("unable to get token: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response body: %s", err)
	}

	var tokens Tokens

	json.Unmarshal(body, &tokens)
	log.Printf("done...")
	return &tokens, nil
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

func findElementBy(wd selenium.WebDriver, by string, selector string) selenium.WebElement {
	elem, err := wd.FindElement(by, selector)
	common.CheckErr(err)
	return elem
}

func sendKeysToElementBy(wd selenium.WebDriver, by string, selector string, keys string) {
	err := findElementBy(wd, by, selector).SendKeys(keys)
	common.CheckErr(err)
}

func sendKeysToElement(element selenium.WebElement, keys string) {
	err := element.SendKeys(keys)
	common.CheckErr(err)
}

func submitElement(element selenium.WebElement) {
	err := element.Submit()
	common.CheckErr(err)
}

func initSelenium() (selenium.WebDriver, *selenium.Service) {
	chromeDriverPath := common.Getenv("CHROMEDRIVER_BINARY", "chromedriver")
	chromeDriverPort := common.Getenv("CHROMEDRIVER_PORT", "9515")

	port, err := strconv.Atoi(chromeDriverPort)
	common.CheckErr(err)

	service, err := selenium.NewChromeDriverService(chromeDriverPath, port)
	common.CheckErr(err)

	chromeOptions := map[string]interface{}{
		"args": []string{
			"--headless",
			"--window-size=1920,1080",
			"--window-position=0,0",
		},
	}

	caps := selenium.Capabilities{
		"browserName":   "chrome",
		"chromeOptions": chromeOptions,
	}

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	common.CheckErr(err)
	return wd, service
}
