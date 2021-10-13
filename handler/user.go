package handler

import (
	"banking_crowd/auth"
	"banking_crowd/helper"
	"banking_crowd/models/tutors"
	"banking_crowd/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	tutorService service.TutorService
	authService  auth.Service
}

func NewUserHandler(tutorService service.TutorService, authService auth.Service) *userHandler {
	return &userHandler{tutorService, authService}
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
		newUser, err := h.tutorService.RegisterTutor(input)
		if err != nil {
			response := helper.APIResponse("Create account failed", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		token, err := h.authService.GenerateToken(newUser.ID)
		if err != nil {
			response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		formatter := tutors.FormatTutor(newUser, token)
		response := helper.APIResponse("Accout has been register", http.StatusOK, "success", formatter)
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
		loggedinUser, err := h.tutorService.Login(input)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			response := helper.APIResponse(err.Error(), http.StatusUnprocessableEntity, "error", errorMessage)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		token, err := h.authService.GenerateToken(loggedinUser.ID)
		if err != nil {
			response := helper.APIResponse(err.Error(), http.StatusBadRequest, "errors", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		formatter := tutors.FormatTutor(loggedinUser, token)

		response := helper.APIResponse("Succesfully Loggedin", http.StatusOK, "success", formatter)

		c.JSON(http.StatusOK, response)

	}

}
