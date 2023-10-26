package service

import (
	"context"
	"github.com/masnann/plant_care/features/assistant/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssistantService_Assistant(t *testing.T) {
	service := NewAssistantService()

	t.Run("Error Case", func(t *testing.T) {
		message := "What's the weather like today?"

		ctx := context.TODO()

		req := domain.MessageRequest{Message: message}
		
		result, err := service.Assistant(ctx, req)

		assert.Error(t, err)
		assert.Empty(t, result)
	})
}
