package handler

import (
	"banking_crowd/auth"
	"banking_crowd/helper"
	"banking_crowd/models/tutors"
	"banking_crowd/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type tutorHandler struct {
	tutorService service.TutorService
	authService  auth.Service
}

func NewTutorHandler(tutorService service.TutorService, authService auth.Service) *tutorHandler {
	return &tutorHandler{tutorService, authService}
}

func (h *tutorHandler) FetchTutor(c *gin.Context) {
	currentTutor := c.MustGet("currentTutor").(tutors.Tutor)
	tutorID := currentTutor.ID
	tutor, err := h.tutorService.GetTutorByID(tutorID)
	if err != nil || tutor.ID < 1 {
		message := "Failed to get tutor"
		if tutor.ID < 1 {
			message = "tutor not found"
		}

		response := helper.APIResponse(message, http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.APIResponse("Success to update tutor", http.StatusOK, "success", tutors.FormatInfoTutor(tutor))
	c.JSON(http.StatusOK, response)
}

func (h *tutorHandler) UpdateTutor(c *gin.Context) {
	var inputID tutors.GetTutorInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update tutor", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	tutor, err := h.tutorService.GetTutorByID(inputID.ID)
	if err != nil || tutor.ID < 1 {
		message := "Failed to get tutor"
		if tutor.ID < 1 {
			message = "tutor not found"
		}

		response := helper.APIResponse(message, http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	var inputData tutors.CreateTutorInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update tutor", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateTutor, err := h.tutorService.UpdateTutor(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update tutor", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update tutor", http.StatusOK, "success", tutors.FormatInfoTutor(updateTutor))
	c.JSON(http.StatusOK, response)

}
