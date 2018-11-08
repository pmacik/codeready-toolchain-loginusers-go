package config

// SetDefaults sets default values for configuration
func SetDefaults(cfg *Configuration) {
	cfg.Auth = AuthConfiguration{
		OAuth2: OAuth2Configuration{
			ClientID: "740650a2-9c44-4db5-b067-a3d1b2cd2d01",
		},
		ServerAddress: "http://localhost:8089",
	}

	cfg.Chromedriver = ChromedriverConfiguration{
		Binary: "chromedriver",
		Port:   9515,
	}

	cfg.Users = UsersConfiguration{
		CredentialsFile: "users.properties",
		Tokens: TokensConfiguration{
			File:            "users.tokens",
			IncludeUsername: false,
		},
		MaxUsers: -1,
	}
}
