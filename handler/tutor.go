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

func (h *tutorHandler) UpdateTutor(c *gin.Context) {
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

	response := helper.APIResponse("Success to update tutor", http.StatusOK, "success", tutor)
	c.JSON(http.StatusOK, response)

}
