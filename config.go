package gw2

import (
	"encoding/json"
	"os"
)

// Config stores the configuration for the application Right now, there are two required objects required:
// - apiKey: <String> - API key generated in the Account page of guildwars2.com
// - baseURL: <String> - baseURL to use for API access. Currently only 'https://api.guildwars2.com/'
type Config struct {
	APIKey  string `json:"apiKey"`
	BaseURL string `json:"baseURL"`
}

// LoadConfig loads the configuration from 'config.json' in the current directory
func LoadConfig(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		return Config{}, err
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	if config.BaseURL == "" {
		config.BaseURL = "https://api.guildwars2.com/"
	}
	return config, nil
}
