package handler

import (
	"banking_crowd/helper"
	myclasses "banking_crowd/models/MyClasses"
	"banking_crowd/models/learners"
	"banking_crowd/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type myClassHandler struct {
	myClassService service.MyClassService
	classService   service.ClassService
}

func NewMyClassHandler(myClassService service.MyClassService, classService service.ClassService) *myClassHandler {
	return &myClassHandler{myClassService, classService}
}

func (h *myClassHandler) CreateMyClass(c *gin.Context) {
	var input myclasses.CreateMyClassInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create class", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentLearner := c.MustGet("currentLearner").(learners.Learner)
	learnerID := currentLearner.ID

	isExistMyClass, err := h.myClassService.IsExistMyClass(input, learnerID)
	if err != nil {
		response := helper.APIResponse("Failed check my class", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if isExistMyClass.ID != 0 {
		response := helper.APIResponse("Learner already take this class", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newMyClass, err := h.myClassService.CreateClass(input, learnerID)
	if err != nil {
		response := helper.APIResponse("Create my class failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	reponse := helper.APIResponse("Success create my class", http.StatusOK, "success", myclasses.FormatMyClass(newMyClass))
	c.JSON(http.StatusOK, reponse)

}

func (h *myClassHandler) GetAllMyClass(c *gin.Context) {
	currentLearner := c.MustGet("currentLearner").(learners.Learner)
	learnerID := currentLearner.ID

	myClass, err := h.myClassService.GetAllMyClass(learnerID)
	if err != nil {
		response := helper.APIResponse("Failed to get my course", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	if len(myClass) == 0 {
		response := helper.APIResponse("you don't have a course", http.StatusOK, "success", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("Success to get detail my course", http.StatusOK, "success", myClass)
	c.JSON(http.StatusOK, response)
}
