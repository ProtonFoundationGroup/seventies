package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func parseYaml(filename string, config *Config) error {
	fullpath, err := filepath.Abs(filename)
	if err != nil {
		return err
	}

	if _, err := os.Stat(fullpath); os.IsNotExist(err) {
		return err
	}

	configFile, err := os.Open(fullpath)
	if err != nil {
		return err
	}
	defer configFile.Close()

	if err = yaml.NewDecoder(configFile).Decode(&config); err != nil {
		return err
	}

	return nil
}
