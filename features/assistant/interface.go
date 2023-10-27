package assistant

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/masnann/plant_care/features/assistant/domain"
)

type ServiceAssistantInterface interface {
	Assistant(ctx context.Context, req domain.MessageRequest) (string, error)
}

type HandlerAssistantInterface interface {
	Assistant() echo.HandlerFunc
}
