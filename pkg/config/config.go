package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"strings"
)

type Config struct {
	Port  uint        `yaml:"port"`
	Redis RedisConfig `yaml:"redis,omitempty"`
	Mongo MongoConfig `yaml:"mongo,omitempty"`
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

var ConfigData = Config{
	Mongo: MongoConfig{},
}

func NewConfig(configStr string) (*Config, error) {
	ConfigData := &Config{
		Port:  7771,
		Redis: RedisConfig{},
	}

	if configStr != "" {
		decoder := yaml.NewDecoder(strings.NewReader(configStr))
		if err := decoder.Decode(ConfigData); err != nil {
			return nil, fmt.Errorf("could not parse config: %v", err)
		}
	}

	return ConfigData, nil
}
