package main

import (
	"content-updater/config"
	"content-updater/infrastructure/externalAPI"
	"content-updater/usecase"
	"fmt"
)

const (
	configFilePath = "./config.yml"
)

func init() {
	config.InitConfig(configFilePath)
	externalAPI.InitHatenaEnv()
	externalAPI.InitZennEnv()
}

func main() {
	hatenaRepository := externalAPI.NewHatenaRepository()
	zennRepository := externalAPI.NewZennRepository()
	entryUsecase := usecase.NewEntryUsecase(hatenaRepository, zennRepository)
	if err := entryUsecase.UpdateList(); err != nil {
		fmt.Printf("%+v\n", err)
	}
}
