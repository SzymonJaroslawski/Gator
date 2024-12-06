package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/golang/glog"
)

const (
	configFileName   = ".gatorconfig.json"
	configFolderName = "gator"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (*Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return &Config{}, nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return &Config{}, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return &Config{}, err
	}

	glog.Infof("Using config file: %s", configPath)

	return &config, nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	err := c.write()
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) write() error {
	configFile, err := getConfigPath()
	if err != nil {
		return err
	}

	file, err := os.OpenFile(configFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func createConfig() (string, error) {
	userConfigPath, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	destPath := filepath.Join(userConfigPath, configFolderName)

	err = os.Mkdir(destPath, os.FileMode(os.ModeDir))
	if err != nil {
		return "", nil
	}

	destFile := filepath.Join(destPath, configFileName)
	file, err := os.Create(destFile)
	if err != nil {
		return "", err
	}
	defer file.Close()
	glog.Infof("Creating new config at: %s", destFile)

	configJson, err := json.Marshal(Config{})
	if err != nil {
		return "", err
	}

	_, err = file.Write(configJson)
	if err != nil {
		return "", err
	}

	glog.Infof("Created new config at: %s", destFile)

	return destFile, nil
}

func getConfigPath() (string, error) {
	path, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	destPath := filepath.Join(path, configFolderName, configFileName)

	if _, err := os.Open(destPath); err != nil {
		glog.Warningf("No config file found at: %s", destPath)

		destPath, err = createConfig()
		if err != nil {
			return "", err
		}

		return destPath, nil
	}

	return destPath, nil
}
