package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/masnann/plant_care/features/note"
	"github.com/masnann/plant_care/features/note/domain"
	"github.com/masnann/plant_care/features/plant"
	user "github.com/masnann/plant_care/features/user/domain"
	"github.com/masnann/plant_care/utils"
	"github.com/masnann/plant_care/utils/response"
	"github.com/sirupsen/logrus"
	"mime/multipart"
	"net/http"
)

type NoteHandler struct {
	service      note.ServiceNoteInterface
	servicePlant plant.ServicePlantInterface
}

func NewNoteHandler(service note.ServiceNoteInterface, servicePlant plant.ServicePlantInterface) note.HandlerNoteInterface {
	return &NoteHandler{
		service:      service,
		servicePlant: servicePlant,
	}
}

func (h *NoteHandler) InsertNotes() echo.HandlerFunc {
	return func(c echo.Context) error {
		insertRequest := new(domain.InsertNoteRequest)
		if err := c.Bind(insertRequest); err != nil {
			logrus.Error("Error Binding Data", err.Error())
			return response.SendErrorResponse(c, http.StatusBadRequest, "Error Binding Data")
		}

		if err := utils.ValidateStruct(insertRequest); err != nil {
			logrus.Error("Validation failed:", err.Error())
			return response.SendErrorResponse(c, http.StatusBadRequest, "Validation failed: "+err.Error())
		}

		plants, err := h.servicePlant.GetPlantsByID(uint(insertRequest.PlantID))
		if err != nil {
			return response.SendErrorResponse(c, http.StatusNotFound, "Plant not found")
		}
		currentUser := c.Get("CurrentUser").(*user.UserModel)
		insertRequest.UserId = currentUser.ID

		fmt.Println(insertRequest.UserId)
		fmt.Println(currentUser.ID)
		if plants.UserID != currentUser.ID {
			return response.SendErrorResponse(c, http.StatusForbidden, "Permission denied")
		}
		fmt.Println(currentUser.ID)
		newNote := &domain.NoteModel{
			UserID:      currentUser.ID,
			PlantID:     insertRequest.PlantID,
			Title:       insertRequest.Title,
			Description: insertRequest.Description,
		}

		createdNote, err := h.service.InsertNote(newNote)

		if err != nil {
			return response.SendErrorResponse(c, http.StatusInternalServerError, "Internal Server Error")
		}

		return response.SendSuccessResponse(c, "Success", createdNote)
	}
}

func (h *NoteHandler) InsertNotePhoto() echo.HandlerFunc {
	return func(c echo.Context) error {
		insertRequest := new(domain.InsertNotePhotoRequest)
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
		fmt.Println(insertRequest.NoteId)

		currentUser := c.Get("CurrentUser").(*user.UserModel)

		notes, err := h.service.GetNoteByID(insertRequest.NoteId)
		if err != nil {
			logrus.Error("error :", err.Error())
			return response.SendErrorResponse(c, http.StatusNotFound, "Note not found")
		}

		if notes.UserID != currentUser.ID {
			return response.SendErrorResponse(c, http.StatusForbidden, "Permission denied")
		}
		newPhoto := &domain.PhotoModel{
			NoteID:      insertRequest.NoteId,
			URL:         uploadedURL,
			Description: insertRequest.Description,
		}

		createdNote, err := h.service.InsertNotePhoto(newPhoto)

		if err != nil {
			return response.SendErrorResponse(c, http.StatusInternalServerError, "Internal Server Error")
		}

		return response.SendSuccessResponse(c, "Success", createdNote)
	}
}
