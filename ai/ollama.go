package ai

import (
	"context"

	"github.com/ollama/ollama/api"
)

var client *api.Client
var err error

func init() {
	client, err = api.ClientFromEnvironment()

	if err != nil {
		panic(err)
	}
}

func Ask(prompt string) (string, error) {
	req := &api.GenerateRequest{
		Model:  "llama3.2",
		Prompt: prompt,

		// set streaming to false
		Stream: new(bool),
	}

	var fullResponse string

	respFunc := func(resp api.GenerateResponse) error {
		fullResponse += resp.Response
		return nil
	}

	err := client.Generate(context.Background(), req, respFunc)

	if err != nil {
		return "", err
	}

	return fullResponse, nil
}
