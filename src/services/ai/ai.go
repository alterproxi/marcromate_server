package ai

import (
	"context"
	"server/src/config"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type ai struct {
	assistant *openai.Assistant
}

const (
	test = "asst_TkxRzWfTbuvcXu4oe7S5olrM"
)

func AI() (*ai, error) {
	client := openai.NewClient(option.WithAPIKey(config.OpenaiApiKey()))

	assistant, err := client.Beta.Assistants.Get(context.Background(), test)

	if err != nil {
		return nil, err
	}

	return &ai{
		assistant: assistant,
	}, nil

}

func (ai *ai) Message(message string) {

}
