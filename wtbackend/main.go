package main

import (
	"github.com/Rhaqim/wtbackend/router"
	"github.com/Rhaqim/wtbackend/service"
)

func main() {

	translatorService, err := service.NewTranslatorClient("server:50051")
	if err != nil {
		panic(err)
	}

	wsService := service.NewWebsocketService(translatorService)

	router := router.NewRouter(wsService)
	router.Run(":8080")

}
