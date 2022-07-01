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

//TODO: consider how to avoid using panic in init
func init() {
	if err := config.InitConfig(configFilePath); err != nil {
		panic(err)
	}
	if err := externalAPI.InitHatenaEnv(); err != nil {
		panic(err)
	}
	if err := externalAPI.InitZennEnv(); err != nil {
		panic(err)
	}
}

func main() {
	hatenaRepository := externalAPI.NewHatenaRepository()
	zennRepository := externalAPI.NewZennRepository()
	entryUsecase := usecase.NewEntryUsecase(hatenaRepository, zennRepository)
	if err := entryUsecase.UpdateList(); err != nil {
		fmt.Printf("%+v\n", err)
	}
}
