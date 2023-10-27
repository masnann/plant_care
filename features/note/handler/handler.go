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
	"strconv"
)

type NoteHandler struct {
	service      note.ServiceNoteInterface
	servicePlant plant.ServicePlantInterface
}

func (h *NoteHandler) DeleteNotesPhotos() echo.HandlerFunc {
	return func(c echo.Context) error {
		var photoData domain.PhotoModel
		photoID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			logrus.Error("Bad Request: Invalid photo ID", err.Error())
			return response.SendErrorResponse(c, http.StatusBadRequest, "Bad Request: Invalid photo ID")
		}

		photoData.PhotoID = photoID

		currentUser := c.Get("CurrentUser").(*user.UserModel)

		existingPhoto, err := h.service.GetNotePhotoByID(photoID)
		if err != nil {
			logrus.Error("Photo not found", err.Error())
			return response.SendErrorResponse(c, http.StatusNotFound, "Notes photo not found")
		}

		note, err := h.service.GetNoteByID(existingPhoto.NoteID)
		if err != nil {
			return response.SendErrorResponse(c, http.StatusNotFound, "Note not found")
		}

		if currentUser.ID != note.UserID {
			return response.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized: You don't have permission to update this photo")
		}

		err = h.service.DeleteNotesPhotos(photoID)
		if err != nil {
			logrus.Error("Internal server error", err.Error())
			return response.SendErrorResponse(c, http.StatusInternalServerError, "Internal Server Error")
		}

		return response.SendDeleteResponse(c, "Success")
	}
}

func (h *NoteHandler) UpdateNotesPhotos() echo.HandlerFunc {
	return func(c echo.Context) error {
		updateRequest := new(domain.UpdateNotePhotoRequest)
		photoID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			logrus.Error("Bad Request: Invalid photo ID", err.Error())
			return response.SendErrorResponse(c, http.StatusBadRequest, "Bad Request: Invalid photo ID")
		}

		if err := c.Bind(updateRequest); err != nil {
			logrus.Error("Bad Request: Invalid photo data ", err.Error())
			return response.SendErrorResponse(c, http.StatusBadRequest, "Bad Request: Invalid photo data")
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

		existingPhoto, err := h.service.GetNotePhotoByID(photoID)
		if err != nil {
			logrus.Error("Photo not found", err.Error())
			return response.SendErrorResponse(c, http.StatusNotFound, "Notes photo not found")
		}

		note, err := h.service.GetNoteByID(existingPhoto.NoteID)
		if err != nil {
			return response.SendErrorResponse(c, http.StatusNotFound, "Note not found")
		}

		if currentUser.ID != note.UserID {
			return response.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized: You don't have permission to update this photo")
		}

		newData := &domain.PhotoModel{
			PhotoID:     photoID,
			NoteID:      existingPhoto.NoteID,
			Photo:       uploadedURL,
			Description: updateRequest.Description,
		}

		updatedPhoto, err := h.service.UpdateNotesPhotos(newData)
		if err != nil {
			logrus.Error("Internal server error", err.Error())
			return response.SendErrorResponse(c, http.StatusInternalServerError, "Internal Server Error "+err.Error())
		}

		result := domain.NotesPhotoResponse{
			PhotoID:     updatedPhoto.PhotoID,
			NoteID:      updatedPhoto.NoteID,
			Photo:       uploadedURL,
			Description: updatedPhoto.Description,
		}

		return response.SendSuccessResponse(c, "Success", result)
	}

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
		result := &domain.GetNotesResponse{
			ID:          createdNote.ID,
			UserID:      createdNote.UserID,
			PlantID:     createdNote.PlantID,
			Date:        createdNote.Date,
			Title:       createdNote.Title,
			Description: createdNote.Description,
		}
		return response.SendSuccessResponse(c, "Success", result)
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
			Photo:       uploadedURL,
			Description: insertRequest.Description,
		}

		createdNote, err := h.service.InsertNotePhoto(newPhoto)

		if err != nil {
			return response.SendErrorResponse(c, http.StatusInternalServerError, "Internal Server Error")
		}

		return response.SendSuccessResponse(c, "Success", createdNote)
	}
}

func (h *NoteHandler) GetNotesWithPagination() echo.HandlerFunc {
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

		notes, err := h.service.GetNotesWithPagination(currentUser.ID, offset, pageSize)
		if err != nil {
			logrus.Error("Internal server error", err.Error())
			return response.SendErrorResponse(c, http.StatusInternalServerError, "Internal server error")
		}

		var responseNotes []domain.GetResponse
		for _, req := range notes {
			responseNotes = append(responseNotes, domain.GetResponse{
				ID:          req.ID,
				PlantID:     req.PlantID,
				UserID:      req.UserID,
				Date:        req.Date,
				Title:       req.Title,
				Description: req.Description,
				Photos:      req.Photos,
			})
		}
		totalItems, err := h.service.CountNotes(currentUser.ID)
		if err != nil {
			logrus.Error("Failed to count notes", err.Error())
			return response.SendErrorResponse(c, http.StatusInternalServerError, "Internal server error")
		}
		return response.PaginationResponse(c, notes, int(totalItems), page, pageSize, "Success")
	}
}

func (h *NoteHandler) UpdateNotes() echo.HandlerFunc {
	return func(c echo.Context) error {
		updateRequest := new(domain.UpdateNoteRequest)
		noteID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			logrus.Error("Bad Request: Invalid plant ID", err.Error())
			return response.SendErrorResponse(c, http.StatusBadRequest, "Bad Request: Invalid plant ID")
		}

		if err := c.Bind(updateRequest); err != nil {
			logrus.Error("Bad Request: Invalid plant data ", err.Error())
			return response.SendErrorResponse(c, http.StatusBadRequest, "Bad Request: Invalid plant data")
		}

		currentUser := c.Get("CurrentUser").(*user.UserModel)

		existingNote, err := h.service.GetNoteByID(noteID)
		if err != nil {
			logrus.Error("Plant not found", err.Error())
			return response.SendErrorResponse(c, http.StatusNotFound, "Notes not found")
		}

		if currentUser.ID != existingNote.UserID {
			return response.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized: You don't have permission to update this plant")
		}

		newData := &domain.NoteModel{
			ID:          noteID,
			UserID:      currentUser.ID,
			PlantID:     existingNote.PlantID,
			Date:        existingNote.Date,
			Title:       updateRequest.Title,
			Description: updateRequest.Description,
		}

		updatedNote, err := h.service.UpdateNotes(newData)
		if err != nil {
			logrus.Error("Internal server error", err.Error())
			return response.SendErrorResponse(c, http.StatusInternalServerError, "Internal Server Error "+err.Error())
		}

		result := domain.GetNotesResponse{
			UserID:      updatedNote.UserID,
			PlantID:     updatedNote.PlantID,
			Date:        updatedNote.Date,
			Title:       updatedNote.Title,
			Description: updatedNote.Description,
		}

		return response.SendSuccessResponse(c, "Success", result)
	}
}

func (h *NoteHandler) DeleteNotes() echo.HandlerFunc {
	return func(c echo.Context) error {
		var noteData domain.NoteModel
		noteID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			logrus.Error("Bad Request: Invalid note ID", err.Error())
			return response.SendErrorResponse(c, http.StatusBadRequest, "Bad Request: Invalid note ID")
		}

		noteData.ID = noteID

		currentUser := c.Get("CurrentUser").(*user.UserModel)

		existingNote, err := h.service.GetNoteByID(noteID)
		if err != nil {
			logrus.Error("Note not found", err.Error())
			return response.SendErrorResponse(c, http.StatusNotFound, "Note not found")
		}

		if currentUser.ID != existingNote.UserID {
			return response.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized: You don't have permission to delete this note")
		}

		err = h.service.DeleteNotes(noteID)
		if err != nil {
			logrus.Error("Internal server error", err.Error())
			return response.SendErrorResponse(c, http.StatusInternalServerError, "Internal Server Error")
		}

		return response.SendDeleteResponse(c, "Success")
	}
}
