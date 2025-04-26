package app

import (
	"server/src/config"

	"github.com/labstack/echo/v4"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type App struct {
	AI     *openai.Client
	Server *echo.Echo
}

func NewApp() *App {
	ai := openai.NewClient(option.WithAPIKey(config.OpenaiApiKey()))

	server := echo.New()

	return &App{
		AI:     &ai,
		Server: server,
	}

}
