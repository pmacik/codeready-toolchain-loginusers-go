package common

import (
	"log"
	"os"
)

// Getenv returns a value of environment variable, if it exists,
// returns the default value otherwise.
func Getenv(key string, defaultValue string) string {
	value, found := os.LookupEnv(key)
	if found {
		return value
	}
	return defaultValue
}

// CheckErr checks for errors and logging it to log as Fatal if not nil
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
