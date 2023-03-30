package config

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func UsersCredentials(cfg *Configuration) ([]string, []string) {
	usersCredentialsFile := cfg.Users.CredentialsFile
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

	return userNames, userPasswords
}
