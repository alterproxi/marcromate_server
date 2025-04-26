package main

import (
	"log"
	api "server/src/api/v1/assistant"
	"server/src/app"
	"server/src/config"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		log.Fatal(err.Error())
		return
	}

	application := app.NewApp()

	assistant_api := api.AssistantApi(application)

	application.Server.GET("/", assistant_api.ShortStory)
	application.Server.POST("/request", assistant_api.Request)

	application.Server.Logger.Fatal(application.Server.Start(":3000"))
}
