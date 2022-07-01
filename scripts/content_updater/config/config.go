package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	NumOfEntry          int    `yaml:"NumOfEntry"`
	EntryListFilePath   string `yaml:"EntryListFilePath"`
	EntryListFileFormat string `yaml:"EntryListFileFormat"`
}

var Conf Config

func InitConfig(configFilePath string) error {
	b, _ := os.ReadFile(configFilePath)
	if err := yaml.Unmarshal(b, &Conf); err != nil {
		return err
	}
	return nil
}
