package chatgpt

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

func SendChatRequest(client *openai.Client, prompt string) (string, error) {
	res, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return res.Choices[0].Message.Content, nil
}
