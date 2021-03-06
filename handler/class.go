package handler

import (
	"banking_crowd/helper"
	"banking_crowd/models/classes"
	"banking_crowd/service"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type classHandler struct {
	classService   service.ClassService
	learnerService service.LearnerService
}

func NewClassHandler(classService service.ClassService, learnerService service.LearnerService) *classHandler {
	return &classHandler{classService, learnerService}
}

func (h *classHandler) CreateClass(c *gin.Context) {
	file, err := c.FormFile("modul")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload file", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Retrieve file information
	extension := filepath.Ext(file.Filename)
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := uuid.New().String() + extension

	path := fmt.Sprintf("storage/files/%s", newFileName)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to file modul", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input classes.CreateClassInput
	err = c.Bind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create class", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newClass, err := h.classService.CreateClass(input, path)
	if err != nil {

		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var pendaftar []string
	response := helper.APIResponse("Success to create class", http.StatusOK, "success", classes.FormatInfoClass(newClass, pendaftar))
	c.JSON(http.StatusOK, response)

}

func (h *classHandler) GetAll(c *gin.Context) {
	tutorID := c.Query("tutor_id")
	err := c.BindQuery(tutorID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create class", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	id, _ := strconv.Atoi(tutorID)
	allclass, err := h.classService.GetAll(id)
	if err != nil {

		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	infoclass := []classes.InfoClassFormatter{}
	for i, class := range allclass {
		var emails []string
		learners, err := h.learnerService.GetLearnerByIdCLass(class.ID)
		if err != nil {
			response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		if learners != nil {
			for _, learner := range learners {
				emails = append(emails, learner.Email)
			}
		} else {
			emails = []string{}
		}

		class := classes.InfoClassFormatter{
			ID:        allclass[i].ID,
			TutorID:   allclass[i].TutorID,
			Topik:     allclass[i].Topik,
			Jenis:     allclass[i].Jenis,
			Judul:     allclass[i].Judul,
			Jadwal:    allclass[i].Jadwal,
			LinkZoom:  allclass[i].LinkZoom,
			Deskripsi: allclass[i].Deskripsi,
			Modul:     allclass[i].Modul,
			Pendaftar: emails,
		}
		infoclass = append(infoclass, class)

	}

	response := helper.APIResponse("Success get all class", http.StatusOK, "success", infoclass)
	c.JSON(http.StatusOK, response)

}
