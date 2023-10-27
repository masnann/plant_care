package handler

import (
	"context" // Import the context package
	"github.com/labstack/echo/v4"
	"github.com/masnann/plant_care/features/assistant"
	"github.com/masnann/plant_care/features/assistant/domain"
	"github.com/masnann/plant_care/utils"
	"github.com/masnann/plant_care/utils/response"
	"net/http"
)

type AssistantHandler struct {
	service assistant.ServiceAssistantInterface
}

func NewAssistantHandler(service assistant.ServiceAssistantInterface) assistant.HandlerAssistantInterface {
	return &AssistantHandler{
		service: service,
	}
}

func (h *AssistantHandler) Assistant() echo.HandlerFunc {
	return func(c echo.Context) error {
		request := new(domain.MessageRequest)
		if err := c.Bind(request); err != nil {
			return response.SendErrorResponse(c, http.StatusBadRequest, "Error Binding Data")
		}
		if err := utils.ValidateStruct(request); err != nil {
			return response.SendErrorResponse(c, http.StatusBadRequest, "Validation failed : "+err.Error())
		}

		ctx := context.TODO()
		reply, err := h.service.Assistant(ctx, *request)

		if err != nil {
			return response.SendErrorResponse(c, http.StatusInternalServerError, "Internal Server Error"+err.Error())
		}
		return response.SendAssistantResponse(c, "Success", reply)

	}
}
