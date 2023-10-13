package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const CONFIG_PATH = "config.yml"

type Config struct {
	Github struct {
		Host         string   `yaml:"host"`
		Graphql      string   `yaml:"graphql"`
		Token        string   `yaml:"token"`
		Owner        string   `yaml:"owner"`
		Repositories []string `yaml:"repositories"`
	} `yaml:"github"`
}

func (c Config) String() string {
	return fmt.Sprintf("Github:\n\tHost: %v\n\tGraphql: %v\n\tToken: %v\n\tOwner: %v", c.Github.Host, c.Github.Graphql, c.Github.Token, c.Github.Owner)
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func readConfig() []byte {
	data, err := os.ReadFile(CONFIG_PATH)
	if err != nil {
		processError(err)
	}
	return data
}

func GetConfig() Config {
	cfg := Config{}
	data := readConfig()
	err := yaml.Unmarshal(data, &cfg)
	if err != nil {
		processError(err)
	}
	return cfg
}
