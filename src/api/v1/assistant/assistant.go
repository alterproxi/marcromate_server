package api

import (
	"fmt"
	"net/http"
	"server/src/app"

	"github.com/labstack/echo/v4"
	"github.com/openai/openai-go"
)

type assistant struct {
	app *app.App
}

func AssistantApi(app *app.App) *assistant {
	return &assistant{
		app: app,
	}
}

func (a *assistant) ShortStory(c echo.Context) error {

	prompt := openai.UserMessage("Short story")

	params := openai.ChatCompletionNewParams{
		Seed:     openai.Int(0),
		Model:    openai.ChatModelGPT4o,
		Messages: []openai.ChatCompletionMessageParamUnion{prompt},
	}

	stream := a.app.AI.Chat.Completions.NewStreaming(
		c.Request().Context(),
		params,
	)

	acc := openai.ChatCompletionAccumulator{}

	// Set headers to make it a streaming response
	c.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
	c.Response().Header().Set("Cache-Control", "no-cache")
	c.Response().Header().Set("Connection", "keep-alive")
	c.Response().WriteHeader(http.StatusOK)

	// Stream loop
	for stream.Next() {
		chunk := stream.Current()
		acc.AddChunk(chunk)

		// If normal content finished
		if _, ok := acc.JustFinishedContent(); ok {
			fmt.Println("Content stream finished")
		}

		if len(chunk.Choices) > 0 {
			text := chunk.Choices[0].Delta.Content

			_, err := c.Response().Write([]byte(text))
			if err != nil {
				fmt.Println("Error writing to client:", err.Error())
				break
			}

			c.Response().Flush()
		}
	}

	// Always close the stream
	stream.Close()

	return nil
}

type request struct {
	Message string `json:"message"`
}

func (a *assistant) Request(c echo.Context) error {

	r := &request{}

	if err := c.Bind(&r); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	prompt := openai.UserMessage(r.Message)

	params := openai.ChatCompletionNewParams{
		Seed:     openai.Int(0),
		Model:    openai.ChatModelGPT4o,
		Messages: []openai.ChatCompletionMessageParamUnion{prompt},
	}

	stream := a.app.AI.Chat.Completions.NewStreaming(
		c.Request().Context(),
		params,
	)

	acc := openai.ChatCompletionAccumulator{}

	// Set headers to make it a streaming response
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().Header().Set("Cache-Control", "no-cache")
	c.Response().Header().Set("Connection", "keep-alive")
	c.Response().WriteHeader(http.StatusOK)

	// Stream loop
	for stream.Next() {
		chunk := stream.Current()
		acc.AddChunk(chunk)

		// If normal content finished
		if _, ok := acc.JustFinishedContent(); ok {
			fmt.Println("Content stream finished")
		}

		if len(chunk.Choices) > 0 {
			text := chunk.Choices[0].Delta.Content

			_, err := c.Response().Write([]byte(text))
			if err != nil {
				fmt.Println("Error writing to client:", err.Error())
				break
			}

			c.Response().Flush()
		}
	}

	// Always close the stream
	stream.Close()

	return nil
}
