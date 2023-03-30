package main

import (
	"log"
	"os"

	"github.com/pmacik/loginusers-go/config"
	"github.com/pmacik/loginusers-go/loginusers"
)

func main() {
	cfg := config.NewConfig("example", "example-config", "yml")

	usernames, passwords := config.UsersCredentials(&cfg)

	username, isSet := os.LookupEnv("USERNAME")
	if !isSet {
		username = usernames[0]
	}
	password, isSet := os.LookupEnv("PASSWORD")
	if !isSet {
		password = passwords[0]
	}

	userTokens, err := loginusers.OAuth2(username, password, cfg)

	if err != nil {
		log.Fatalf("Unable to login user: %s", err)
		return
	}

	log.Printf("Auth: %s", userTokens.AccessToken)
	log.Printf("Refresh: %s", userTokens.RefreshToken)
}
