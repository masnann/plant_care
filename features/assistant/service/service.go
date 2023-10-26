package service

import (
	"context"
	"fmt"
	"github.com/masnann/plant_care/features/assistant"
	"github.com/masnann/plant_care/features/assistant/domain"
	"github.com/sashabaranov/go-openai"
	"os"
)

type AssistantService struct {
	client *openai.Client
}

func NewAssistantService() assistant.ServiceAssistantInterface {
	apiKey := os.Getenv("APIKEY_OPENAI")
	client := openai.NewClient(apiKey)

	return &AssistantService{
		client: client,
	}
}

func (s *AssistantService) Assistant(ctx context.Context, req domain.MessageRequest) (string, error) {
	chatMessage := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: fmt.Sprintf("Jawab pertanyaan ini : %s", req.Message),
	}

	chatReq := openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{chatMessage},
	}

	resp, err := s.client.CreateChatCompletion(ctx, chatReq)
	if err != nil {
		return "", err
	}

	reply := resp.Choices[0].Message.Content
	return reply, nil
}
