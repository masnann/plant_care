package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/masnann/plant_care/features/plant"
	"github.com/masnann/plant_care/features/plant/domain"
	user "github.com/masnann/plant_care/features/user/domain"
	"github.com/masnann/plant_care/utils"
	"github.com/masnann/plant_care/utils/response"
	"github.com/sirupsen/logrus"
	"mime/multipart"
	"net/http"
	"strconv"
)

type PlantHandler struct {
	service plant.ServicePlantInterface
	jwt     utils.JWTInterface
}

func NewPlantHandler(service plant.ServicePlantInterface, jwt utils.JWTInterface) plant.HandlerPlantInterface {
	return &PlantHandler{
		service: service,
		jwt:     jwt,
	}
}

func (h *PlantHandler) SearchPlantsByType() echo.HandlerFunc {
	return func(c echo.Context) error {
		types := c.QueryParam("types")
		if types == "" {
			return response.SendErrorResponse(c, http.StatusBadRequest, "Parameter 'types' is required")
		}
		currentUser := c.Get("CurrentUser").(*user.UserModel)
		if currentUser == nil {
			return response.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized: You don't have permission")
		}

		plants, err := h.service.SearchPlantsByType(currentUser.ID, types)
		if err != nil {
			return response.SendErrorResponse(c, http.StatusInternalServerError, "Failed to search plants by type: "+err.Error())
		}
		var responsePlants []domain.GetResponse
		for _, plants := range plants {
			responsePlants = append(responsePlants, domain.GetResponse{
				ID:    plants.ID,
				Name:  plants.Name,
				Type:  plants.Type,
				Date:  plants.Date,
				Photo: plants.Photo,
			})
		}
		if len(responsePlants) == 0 {
			return response.SendSuccessResponse(c, "Success", nil)
		} else {
			return response.SendSuccessResponse(c, "Success", responsePlants)
		}
	}
}

func (h *PlantHandler) SearchPlantsByName() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.QueryParam("name")
		if name == "" {
			return response.SendErrorResponse(c, http.StatusBadRequest, "Parameter 'name' is required")
		}
		currentUser := c.Get("CurrentUser").(*user.UserModel)
		if currentUser == nil {
			return response.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized: You don't have permission ")
		}

		plants, err := h.service.SearchPlantsByName(currentUser.ID, name)
		if err != nil {
			return response.SendErrorResponse(c, http.StatusInternalServerError, "Failed to search plants by name: "+err.Error())
		}
		var responsePlants []domain.GetResponse
		for _, plants := range plants {
			responsePlants = append(responsePlants, domain.GetResponse{
				ID:    plants.ID,
				Name:  plants.Name,
				Type:  plants.Type,
				Date:  plants.Date,
				Photo: plants.Photo,
			})
		}
		if len(responsePlants) == 0 {
			return response.SendSuccessResponse(c, "Success", nil)
		} else {
			return response.SendSuccessResponse(c, "Success", responsePlants)
		}
	}
}

func (h *PlantHandler) InsertPlants() echo.HandlerFunc {
	return func(c echo.Context) error {
		insertRequest := new(domain.InsertRequest)
		file, err := c.FormFile("photo")
		var uploadedURL string
		if err == nil {
			fileToUpload, err := file.Open()
			if err != nil {
				return response.SendErrorResponse(c, http.StatusInternalServerError, "Failed to open file"+err.Error())
			}
			defer func(fileToUpload multipart.File) {
				err := fileToUpload.Close()
				if err != nil {

				}
			}(fileToUpload)

			uploadedURL, err = utils.ImageUploadHelper(fileToUpload)
			if err != nil {
				return response.SendErrorResponse(c, http.StatusInternalServerError, "Image upload error "+err.Error())
			}
		}

		if err := c.Bind(insertRequest); err != nil {
			logrus.Error("Error Binding Data", err.Error())
			return response.SendErrorResponse(c, http.StatusBadRequest, "Error Binding Data")
		}

		if err := utils.ValidateStruct(insertRequest); err != nil {
			logrus.Error("Validation failed:", err.Error())
			return response.SendErrorResponse(c, http.StatusBadRequest, "Validation failed: "+err.Error())
		}

		currentUser := c.Get("CurrentUser").(*user.UserModel)
		insertRequest.UserID = currentUser.ID

		newPlant := &domain.PlantModel{
			UserID: insertRequest.UserID,
			Name:   insertRequest.Name,
			Type:   insertRequest.Type,
			Photo:  uploadedURL,
		}

		createdPlant, err := h.service.InsertPlants(newPlant)
		if err != nil {
			return response.SendErrorResponse(c, http.StatusInternalServerError, "Internal Server Error")
		}

		result := domain.InsertResponse{
			Name:  createdPlant.Name,
			Type:  createdPlant.Type,
			Date:  createdPlant.Date,
			Photo: createdPlant.Photo,
		}
		return response.SendSuccessResponse(c, "Success", result)
	}
}

func (h *PlantHandler) GetPaginationPlants() echo.HandlerFunc {
	return func(c echo.Context) error {
		currentUser := c.Get("CurrentUser").(*user.UserModel)

		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil {
			logrus.Error("Invalid page parameter", err.Error())
			return response.SendErrorResponse(c, http.StatusBadRequest, "Invalid page parameter ")
		}

		pageSize, err := strconv.Atoi(c.QueryParam("pageSize"))
		if err != nil {
			logrus.Error("Invalid pageSize parameter", err.Error())
			return response.SendErrorResponse(c, http.StatusBadRequest, "Invalid pageSize parameter ")
		}

		offset := (page - 1) * pageSize

		plants, err := h.service.GetPlantsWithPagination(currentUser.ID, offset, pageSize)
		if err != nil {
			logrus.Error("Internal server error", err.Error())
			return response.SendErrorResponse(c, http.StatusInternalServerError, "Internal server error")
		}
		var responsePlants []domain.GetResponse
		for _, plants := range plants {
			responsePlants = append(responsePlants, domain.GetResponse{
				ID:    plants.ID,
				Name:  plants.Name,
				Type:  plants.Type,
				Date:  plants.Date,
				Photo: plants.Photo,
			})
		}
		totalItems, err := h.service.CountPlants(currentUser.ID)
		if err != nil {
			logrus.Error("Failed to count plants", err.Error())
			return response.SendErrorResponse(c, http.StatusInternalServerError, "Internal server error")
		}
		return response.PaginationResponse(c, responsePlants, int(totalItems), page, pageSize, "Success")
	}
}

func (h *PlantHandler) UpdatePlants() echo.HandlerFunc {
	return func(c echo.Context) error {
		updateRequest := new(domain.UpdateRequest)
		plantID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			logrus.Error("Bad Request: Invalid plant ID", err.Error())
			return response.SendErrorResponse(c, http.StatusBadRequest, "Bad Request: Invalid plant ID")
		}

		if err := c.Bind(updateRequest); err != nil {
			logrus.Error("Bad Request: Invalid plant data ", err.Error())
			return response.SendErrorResponse(c, http.StatusBadRequest, "Bad Request: Invalid plant data")
		}

		file, err := c.FormFile("photo")
		var uploadedURL string
		if err == nil {
			fileToUpload, err := file.Open()
			if err != nil {
				return response.SendErrorResponse(c, http.StatusInternalServerError, "Failed to open file "+err.Error())
			}
			defer func(fileToUpload multipart.File) {
				err := fileToUpload.Close()
				if err != nil {

				}
			}(fileToUpload)

			uploadedURL, err = utils.ImageUploadHelper(fileToUpload)
			if err != nil {
				return response.SendErrorResponse(c, http.StatusInternalServerError, "Image upload error "+err.Error())
			}
		}

		updateRequest.Photo = uploadedURL

		currentUser := c.Get("CurrentUser").(*user.UserModel)

		existingPlant, err := h.service.GetPlantsByID(uint(plantID))
		if err != nil {
			logrus.Error("Plant not found", err.Error())
			return response.SendErrorResponse(c, http.StatusNotFound, "Plant not found")
		}

		if currentUser.ID != existingPlant.UserID {
			return response.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized: You don't have permission to update this plant")
		}

		newData := &domain.PlantModel{
			ID:    plantID,
			Name:  updateRequest.Name,
			Type:  updateRequest.Type,
			Photo: updateRequest.Photo,
		}

		updatedPlant, err := h.service.UpdatePlants(newData)
		if err != nil {
			logrus.Error("Internal server error", err.Error())
			return response.SendErrorResponse(c, http.StatusInternalServerError, "Internal Server Error "+err.Error())
		}

		result := domain.UpdateResponse{
			Name:  updatedPlant.Name,
			Type:  updatedPlant.Type,
			Date:  updatedPlant.Date,
			Photo: updatedPlant.Photo,
		}

		return response.SendSuccessResponse(c, "Success", result)
	}
}

func (h *PlantHandler) DeletePlants() echo.HandlerFunc {
	return func(c echo.Context) error {
		var plantData domain.PlantModel
		plantID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			logrus.Error("Bad Request: Invalid plant ID", err.Error())
			return response.SendErrorResponse(c, http.StatusBadRequest, "Bad Request: Invalid plant ID")
		}

		plantData.ID = plantID

		currentUser := c.Get("CurrentUser").(*user.UserModel)

		existingPlant, err := h.service.GetPlantsByID(uint(plantID))
		if err != nil {
			logrus.Error("Plant not found", err.Error())
			return response.SendErrorResponse(c, http.StatusNotFound, "Plant not found")
		}

		if currentUser.ID != existingPlant.UserID {
			return response.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized: You don't have permission to update this plant")
		}

		err = h.service.DeletePlants(plantID)
		if err != nil {
			logrus.Error("Internal server error", err.Error())
			return response.SendErrorResponse(c, http.StatusInternalServerError, "Internal Server Error")
		}

		return response.SendDeleteResponse(c, "Success")
	}
}
