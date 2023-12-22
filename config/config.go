package config

import (
	"github.com/jbillote/YAB/util/logger"

	"gopkg.in/yaml.v2"

	"os"
)

type Config struct {
	BotToken       string   `yaml:"Token"`
	APIHost        string   `yaml:"Host"`
	CommandPrefix  string   `yaml:"Prefix"`
	Statuses       []string `yaml:"Statuses"`
	StatusInterval int64    `yaml:"StatusInterval"`
}

var log = logger.GetLogger("Config")

var config *Config = nil

func GetConfig() *Config {
	return config
}

func CreateConfig(file *os.File) (*Config, error) {
	decoder := yaml.NewDecoder(file)
	err := decoder.Decode(config)
	if err != nil {
		log.Fatal("Unable to parse config file")
		log.Fatal(err)
		return nil, err
	}

	return config, nil
}

func SetConfig(c *Config) {
	config = c
}
