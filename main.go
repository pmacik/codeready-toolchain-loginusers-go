package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pmacik/loginusers-go/common"
	"github.com/pmacik/loginusers-go/config"
	"github.com/pmacik/loginusers-go/loginusers"
)

func main() {
	cfg := config.Config()

	authServerAddress := cfg.Auth.ServerAddress
	usersCredentialsFile := cfg.Users.CredentialsFile
	userTokensFile := cfg.Users.Tokens.File
	userTokensIncludeUsername := cfg.Users.Tokens.IncludeUsername

	maxUsers := cfg.Users.MaxUsers

	i=5

	log.SetOutput(os.Stdout)

	ufile, err := os.Open(usersCredentialsFile)
	if err != nil {
		log.Fatal(err)
	}
	defer ufile.Close()

	scanner := bufio.NewScanner(ufile)

	var userNames []string
	var userPasswords []string

	for scanner.Scan() {
		line := scanner.Text()
		credentials := strings.Split(line, "=")
		userNames = append(userNames, credentials[0])
		userPasswords = append(userPasswords, credentials[1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat(userTokensFile); os.IsExist(err) {
		os.Remove(userTokensFile)
	}
	tfile, err := os.Create(userTokensFile)
	common.CheckErr(err)
	defer tfile.Close()

	w := bufio.NewWriter(tfile)
	defer w.Flush()
	for index, userName := range userNames {
		if maxUsers > 0 && index >= maxUsers {
			break
		}
		log.Printf("Loggin user %s via %s", userName, authServerAddress)
		tokens, err := loginusers.OAuth2(userName, userPasswords[index], cfg)
		common.CheckErr(err)
		tokenLine := fmt.Sprintf("%s;%s", tokens.AccessToken, tokens.RefreshToken)
		if userTokensIncludeUsername {
			tokenLine = fmt.Sprintf("%s;%s", tokenLine, userName)
		}
		//write tokens to user.tokens file
		_, err = w.WriteString(fmt.Sprintf("%s\n", tokenLine))
		common.CheckErr(err)
	}
}
