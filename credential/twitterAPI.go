package credential

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type TwitterCred struct {
	TwitterAPIKeys TwitterAPIKey `yaml:"twitterAPIKeys"`
}

type TwitterAPIKey struct {
	ConsumerKey       string `yaml:"consumerKey"`
	ConsumerSecret    string `yaml:"consumerSecret"`
	AccessToken       string `yaml:"accessToken"`
	AccessTokenSecret string `yaml:"accessTokenSecret"`
	CallbackURL       string `yaml:"callbackURL"`
}

func GetTwitterAPIKeys() *TwitterCred {
	credFile, err := os.Open("/usr/local/bin/config.yaml")
	if err != nil {
		fmt.Println("Can't find a config.yaml file!")
	}

	yamlFile, err := ioutil.ReadAll(credFile)

	twitterCred := TwitterCred{}

	if err = yaml.Unmarshal([]byte(yamlFile), &twitterCred); err != nil {
		fmt.Println("Can't unmarshal a config.yaml file!")
	}
	return &twitterCred
}
