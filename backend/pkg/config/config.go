package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Server struct {
		Address string `yaml:"address"`
	} `yaml:"server"`
	LogLevel string `yaml:"log_level"`
	DB       struct {
		Address string `yaml:"address"`
	} `yaml:"db"`
}

func New(configPath string) (*Config, error) {

	config := &Config{
		Server: struct {
			Address string `yaml:"address"`
		}{
			Address: "127.0.0.1:8080",
		},
		LogLevel: "debug",
	}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
