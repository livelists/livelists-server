package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"strings"
)

type ConfStore struct {
	Conf *Config
}

var confStore ConfStore

type Config struct {
	ClientPort uint        `yaml:"clientPort"`
	AdminPort  uint        `yaml:"adminPort"`
	ApiKey     string      `yaml:"api_key"`
	SecretKey  string      `yaml:"secret_key"`
	Mongo      MongoConfig `yaml:"mongo,omitempty"`
}

type RedisConfig struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type MongoConfig struct {
	URI string `yaml:"uri"`
}

func NewConfig(configStr string) (*Config, error) {
	configData := &Config{
		AdminPort:  7772,
		ClientPort: 7771,
	}

	if configStr != "" {
		decoder := yaml.NewDecoder(strings.NewReader(configStr))
		if err := decoder.Decode(configData); err != nil {
			return nil, fmt.Errorf("could not parse config: %v", err)
		}
	}

	confStore.Conf = configData

	return configData, nil
}

func GetConfig() Config {
	if confStore.Conf == nil {
		confStore.Conf = &Config{}
	}
	return *confStore.Conf
}
