package assistant

import (
	"context"
	"fmt"

	"github.com/openai/openai-go"
)

type Assistant struct {
	Client *openai.Client
}

func (assistant *Assistant) Test(context context.Context) {
	params := openai.BetaAssistantListParams{}
	result, err := assistant.Client.Beta.Assistants.List(context, params)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(result)

}
