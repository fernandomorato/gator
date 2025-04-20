package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	return write(*c)
}

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	file, err := os.Open(configFilePath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()
	config := Config{}
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func write(config Config) error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	file, err := os.Create(configFilePath)
	if err != nil {
		return err
	}
	err = json.NewEncoder(file).Encode(config)
	if err != nil {
		return err
	}
	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, configFileName), nil
}
