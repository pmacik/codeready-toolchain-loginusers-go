package main

import (
	"log"

	"github.com/pmacik/loginusers-go/config"
	"github.com/pmacik/loginusers-go/loginusers"
)

func main() {
	cfg := config.DefaultConfig()
	cfg.Auth.ServerAddress = "http://localhost:8089"

	userTokens, err := loginusers.OAuth2("username", "password", cfg)

	if err != nil {
		log.Fatalf("Unable to login user: %s", err)
		return
	}

	log.Printf("Auth: %s", userTokens.AccessToken)
	log.Printf("Refresh: %s", userTokens.RefreshToken)
}
