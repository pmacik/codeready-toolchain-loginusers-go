# Login Openshift.io users

An utility to login Openshift.io users and get auth and refresh tokens.

## Prerequisities

Chrome or [Chromium browser](https://www.chromium.org/Home) with headless feature and [Chromedriver](https://sites.google.com/a/chromium.org/chromedriver/) needs to be installed where it is run (for Fedora/RHEL/CentOS):

```shell
sudo yum install chromium chromium-headless chromedriver
```

## Usage

To run, provide a line separated list of users ("user=password") in the property file defined by the `users.credentialsFile` variable in the `config.yml` (default file is `users.properties`)

### Configuration via `config.xml`

* `auth.serverAddress` = server of Auth Services (default `https://localhost:8089`).
* `auth.oauth2.clientID` = OAuth2 protocol client id (default `740650a2-9c44-4db5-b067-a3d1b2cd2d01`).
* `users.credentialsFile` = a file containing a line separated list of users in a form of `user=password` (default `users.properties`).
* `users.tokens.file` = an output file where the generated auth and refresh tokens were written after succesfull login of each user (default `users.tokens`).
* `users.tokens.includeUsername` = "`true` if username is to be included in the output (default `talse`).
* `users.maxUsers` = A maximal number of users taken from the `users.credentialsFile` (default `-1` means unlimited).

### Run standalone

```shell
go run main.go
```

### Use as Go library

```go
package main

import (
    "log"

    "github.com/pmacik/loginusers-go/config"
    "github.com/pmacik/loginusers-go/loginuser"
)

func main(){
    cfg := config.DefaultConfig()
    cfg..Auth.ServerAddress = "http://localhost:8089"

    userTokens, err := loginusers.OAuth2("username", "password", cfg)

    if err != nil {
        log.Fatalf("Unable to login user: %s", err)
        return
    }

    log.Printf("Auth: %s", userTokens.AccessToken)
    log.Printf("Refresh: %s", userTokens.RefreshToken)
}
```