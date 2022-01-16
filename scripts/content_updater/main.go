package main

import (
	"content-updater/config"
	"content-updater/infrastructure/externalAPI"
	"content-updater/usecase"
)

const (
	configFilePath = "./config.yml"
)

func init() {
	config.InitConfig(configFilePath)
	externalAPI.InitHatenaEnv()
}

func main() {
	hatenaRepository := externalAPI.NewHatenaRepository()
	entryUsecase := usecase.NewEntryUsecase(hatenaRepository)
	entryUsecase.UpdateList()
}
