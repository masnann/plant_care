package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/masnann/plant_care/features/guide"
	"github.com/masnann/plant_care/features/guide/domain"
	"github.com/masnann/plant_care/utils/response"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type GuideHandler struct {
	service guide.ServiceGuideInterface
}

func NewGuideHandler(service guide.ServiceGuideInterface) guide.HandlerGuideInterface {
	return &GuideHandler{
		service: service,
	}
}
func (h *GuideHandler) SearchGuideByName() echo.HandlerFunc {
	return func(c echo.Context) error {
		guideName := c.QueryParam("name")
		if guideName == "" {
			return response.SendErrorResponse(c, http.StatusBadRequest, "Parameter 'name' is required")
		}
		guides, err := h.service.SearchGuideByName(guideName)
		if err != nil {
			logrus.Error("Failed to search guide by name: ", err.Error())
			return response.SendErrorResponse(c, http.StatusInternalServerError, "Status internal server error: "+err.Error())
		}
		var responseGuides []domain.GetGuideResponse
		for _, res := range guides {
			responseGuides = append(responseGuides, domain.GetGuideResponse{
				ID:          res.ID,
				Title:       res.Title,
				Description: res.Description,
				Date:        res.Date,
				Photo:       res.Photo,
			})
		}
		return response.SendSuccessResponse(c, "Success", responseGuides)

	}
}

func (h *GuideHandler) GetGuidesById() echo.HandlerFunc {
	return func(c echo.Context) error {
		guideIDString := c.Param("id")
		guideID, err := strconv.ParseUint(guideIDString, 10, 64)
		if err != nil {
			return response.SendErrorResponse(c, http.StatusBadRequest, "Parameter 'id' is required"+err.Error())
		}
		guides, err := h.service.GetByIdGuides(guideID)
		if err != nil {
			logrus.Error("Failed to get guide by Id: ", err.Error())
			return response.SendErrorResponse(c, http.StatusInternalServerError, "Status internal server error: "+err.Error())
		}
		responseGuides := domain.GetGuideResponse{
			ID:          guides.ID,
			Title:       guides.Title,
			Description: guides.Description,
			Date:        guides.Date,
			Photo:       guides.Photo,
		}
		return response.SendSuccessResponse(c, "Success", responseGuides)

	}
}

func (h *GuideHandler) GetGuidesWithPagination() echo.HandlerFunc {
	return func(c echo.Context) error {
		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil || page <= 0 {
			return response.SendErrorResponse(c, http.StatusBadRequest, "Parameter 'page' is required: "+err.Error())
		}

		pageSize, err := strconv.Atoi(c.QueryParam("pageSize"))
		if err != nil || pageSize <= 0 {
			return response.SendErrorResponse(c, http.StatusBadRequest, "Parameter 'pageSize' is required: "+err.Error())
		}

		guides, err := h.service.GetGuidesWithPagination(page, pageSize)
		if err != nil {
			logrus.Error("Failed to search guide by name: ", err.Error())
			return response.SendErrorResponse(c, http.StatusInternalServerError, "Status internal server error: "+err.Error())
		}
		var responseGuide []domain.GetGuideResponse
		for _, res := range guides {
			responseGuide = append(responseGuide, domain.GetGuideResponse{
				ID:          res.ID,
				Title:       res.Title,
				Description: res.Description,
				Date:        res.Date,
				Photo:       res.Photo,
			})
		}

		totalItems, err := h.service.CountGuides()
		if err != nil {
			logrus.Error("Failed to count plants", err.Error())
			return response.SendErrorResponse(c, http.StatusInternalServerError, "Status internal server error: "+err.Error())
		}
		return response.PaginationResponse(c, responseGuide, int(totalItems), page, pageSize, "Success")
	}
}
