package main

import (
	"content-updater/config"
	"content-updater/infrastructure/externalAPI"
	"content-updater/usecase"
)

func init() {
	config.InitConfig()
	externalAPI.InitHatenaEnv()
}

func main() {
	hatenaRepository := externalAPI.NewHatenaRepository()
	entryUsecase := usecase.NewEntryUsecase(hatenaRepository)
	entryUsecase.UpdateEntryList()
}
