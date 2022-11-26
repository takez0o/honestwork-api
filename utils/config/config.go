package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	API struct {
		Auth struct {
			Port string `yaml:"port"`
			ID   int    `yaml:"id"`
		}
		Broker struct {
			Port string `yaml:"port"`
			ID   int    `yaml:"id"`
		}
		Messages struct {
			Port string `yaml:"port"`
			ID   int    `yaml:"id"`
		}
		Posts struct {
			Port string `yaml:"port"`
			ID   int    `yaml:"id"`
		}
	}

	DB struct {
		Users struct {
			ID int `yaml:"id"`
		}
		Posts struct {
			ID int `yaml:"id"`
		}
	}
}

// ParseConfig
func ParseConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var conf Config
	if err := yaml.Unmarshal(data, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
