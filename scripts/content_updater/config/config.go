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

func InitConfig() {
	b, _ := os.ReadFile("../config.yml")
	yaml.Unmarshal(b, &Conf)
}
