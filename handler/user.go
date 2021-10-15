package handler

import (
	"banking_crowd/auth"
	"banking_crowd/helper"
	"banking_crowd/models/learners"
	"banking_crowd/models/tutors"
	"banking_crowd/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	tutorService   service.TutorService
	learnerService service.LearnerService
	authService    auth.Service
}

func NewUserHandler(tutorService service.TutorService, learnerService service.LearnerService, authService auth.Service) *userHandler {
	return &userHandler{tutorService, learnerService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {

	var input tutors.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Accout Register failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	role, _ := strconv.Atoi(input.RegisterAs)
	if role == 1 {
		newTutor, err := h.tutorService.RegisterTutor(input)
		if err != nil {
			response := helper.APIResponse("Create account tutor failed", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		token, err := h.authService.GenerateToken(newTutor.ID, role)
		if err != nil {
			response := helper.APIResponse("Register account tutor failed", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		formatter := tutors.FormatTutor(newTutor, token)
		response := helper.APIResponse("Accout has been register as tutor", http.StatusOK, "success", formatter)
		c.JSON(http.StatusOK, response)
	} else {
		newLearner, err := h.learnerService.RegisterLearner(input)
		if err != nil {
			response := helper.APIResponse("Create account learner failed", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		token, err := h.authService.GenerateToken(newLearner.ID, role)
		if err != nil {
			response := helper.APIResponse("Register account failed learner", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		formatter := learners.Formatlearner(newLearner, token)
		response := helper.APIResponse("Accout has been register as learner", http.StatusOK, "success", formatter)
		c.JSON(http.StatusOK, response)
	}

}

func (h *userHandler) Login(c *gin.Context) {

	var input tutors.LogisUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	role, _ := strconv.Atoi(input.LoginAs)
	if role == 1 {
		loggedinTutor, err := h.tutorService.Login(input)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			response := helper.APIResponse(err.Error(), http.StatusUnprocessableEntity, "error", errorMessage)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		token, err := h.authService.GenerateToken(loggedinTutor.ID, role)
		if err != nil {
			response := helper.APIResponse(err.Error(), http.StatusBadRequest, "errors", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		formatter := tutors.FormatTutor(loggedinTutor, token)

		response := helper.APIResponse("Succesfully Loggedin", http.StatusOK, "success", formatter)

		c.JSON(http.StatusOK, response)

	} else {
		loggedinLearner, err := h.learnerService.LoginLearner(input)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			response := helper.APIResponse(err.Error(), http.StatusUnprocessableEntity, "error", errorMessage)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		token, err := h.authService.GenerateToken(loggedinLearner.ID, role)
		if err != nil {
			response := helper.APIResponse(err.Error(), http.StatusBadRequest, "errors", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		formatter := learners.Formatlearner(loggedinLearner, token)

		response := helper.APIResponse("Succesfully Loggedin", http.StatusOK, "success", formatter)

		c.JSON(http.StatusOK, response)
	}

}
