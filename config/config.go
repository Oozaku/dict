/*
Package config is responsible to get user's configuration located at 
$HOME/.config/dict/dict.json.
*/
package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// GetConfiguration loads the configuration file from user and returns the
// struct Config and nil if it was possible to load the file without failures.
// If an error occurs, it returns an empty struct followed by error.
func GetConfiguration() (Config, error) {

  // Get home location
	home, err := os.UserHomeDir()
	if err != nil {
		log.Println(err)
		return Config{}, err
	}

  // Construct path where config file should be located
	configPath := filepath.Join(home, ".config", "dict", "dict.json")

  // Read configuration file
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Println(err)
		return Config{}, err
	}

  // Build struct with user's values
	config := Config{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Println(err)
		return Config{}, err
	}

  // Configuration file loaded successfully: return config and nil
	return config, nil
}
