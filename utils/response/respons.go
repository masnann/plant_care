package response

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type AssistantResponse struct {
	Message string      `json:"message"`
	Reply   interface{} `json:"reply"`
}

func SendErrorResponse(c echo.Context, status int, message string) error {
	return c.JSON(status, ErrorResponse{
		Message: message,
	})
}

func SendSuccessResponse(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusOK, SuccessResponse{
		Message: message,
		Data:    data,
	})
}

func SendDeleteResponse(c echo.Context, message string) error {
	return c.JSON(http.StatusOK, ErrorResponse{
		Message: message,
	})
}

func SendAssistantResponse(c echo.Context, message string, reply interface{}) error {
	return c.JSON(http.StatusOK, AssistantResponse{
		Message: message,
		Reply:   reply,
	})
}

func PaginationResponse(c echo.Context, data interface{}, totalItems, page, pageSize int, message string) error {
	pagination := map[string]interface{}{
		"totalItems": totalItems,
		"page":       page,
		"pageSize":   pageSize,
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: message,
		Data:    map[string]interface{}{"items": data, "pagination": pagination},
	})
}
