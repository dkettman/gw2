package gw2

import (
	"encoding/json"
	"log"
	"os"
)

// LoadConfig loads the configuration from 'config.json' in the current directory
func LoadConfig(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		log.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
