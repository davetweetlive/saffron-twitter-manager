package config

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/spf13/viper"
)

// Credentials ...
type Credentials struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

// GetCredentials ...
func (c Credentials) GetCredentials() Credentials {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath("./config")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Set undefined variables
	viper.SetDefault("database.dbname", "test_db")

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	// Reading variables using the model
	// fmt.Println("Reading variables using the model..")
	// fmt.Println("Database is\t", configuration.KeysAndTokens.ConsumerKey)
	// fmt.Println("Port is\t\t", configuration.Server.Port)
	cd := Credentials{
		ConsumerKey:    configuration.KeysAndTokens.ConsumerKey,
		ConsumerSecret: configuration.KeysAndTokens.ConsumerSecret,
		AccessToken:    configuration.KeysAndTokens.AccessToken,
		AccessSecret:   configuration.KeysAndTokens.AccessSecret,
	}
	return cd
}

func TwitterClient() *twitter.Client {
	someToken := Credentials{}
	config := oauth1.NewConfig(someToken.GetCredentials().ConsumerKey, someToken.GetCredentials().ConsumerSecret)
	token := oauth1.NewToken(someToken.GetCredentials().AccessToken, someToken.GetCredentials().AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)
	return client
}
