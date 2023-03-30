package config

// AuthConfiguration handles configuration for Auth service
type AuthConfiguration struct {
	OAuth2        OAuth2Configuration
	Tokens        TokensConfiguration
	ServerAddress string
	RedirectURI   string
	Path          string
}
