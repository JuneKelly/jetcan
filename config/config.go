package config

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	ServiceUri	string `json:"serviceUri"`
	StorageDir	string `json:"storageDir"`
}

/*
Get the location of the jetcan config file,
falling back to using "~/.jetcan"
*/
func findConfigFile() (string, error) {
	configFile := os.Getenv("JETCAN_CONFIG")
	if configFile == "" {
		currentUser, err := user.Current()
		if err != nil {
			return "", err
		}
		configFile = filepath.Join(currentUser.HomeDir, ".jetcan")
	}
	return configFile, nil
}

/*
Try to load configuration from json config file
*/
func Load() (*Config, error) {
	var cfg	Config

	configFile, err := findConfigFile()
	if err != nil {
		return nil, err
	}

	f, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(f, &cfg)
	if err != nil {
		fmt.Println("Error loading config file")
		return nil, err
	}
	return &cfg, nil
}
