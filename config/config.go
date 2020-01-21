package config

// Configurations exported
type Configurations struct {
	Server        ServerConfigurations
	KeysAndTokens KeysAndTokensConf
	// EXAMPLE_PATH  string
	// EXAMPLE_VAR   string
}

// ServerConfigurations exported
type ServerConfigurations struct {
	Port int
}

// KeysAndTokensConf exported
type KeysAndTokensConf struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}
