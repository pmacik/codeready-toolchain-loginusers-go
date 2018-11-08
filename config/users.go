package config

// UsersConfiguration handles configuration related to users
type UsersConfiguration struct {
	CredentialsFile string
	Tokens          TokensConfiguration
	MaxUsers        int
}
