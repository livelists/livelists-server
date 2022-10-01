package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"strings"
)

type Config struct {
	Port  uint        `yaml:"port"`
	Redis RedisConfig `yaml:"redis,omitempty"`
}

type RedisConfig struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

func NewConfig(configStr string) (*Config, error) {
	conf := &Config{
		Port:  7771,
		Redis: RedisConfig{},
	}

	if configStr != "" {
		decoder := yaml.NewDecoder(strings.NewReader(configStr))
		if err := decoder.Decode(conf); err != nil {
			return nil, fmt.Errorf("could not parse config: %v", err)
		}
	}

	return conf, nil
}
