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

func InitConfig(configFilePath string) {
	b, _ := os.ReadFile(configFilePath)
	yaml.Unmarshal(b, &Conf)
}
