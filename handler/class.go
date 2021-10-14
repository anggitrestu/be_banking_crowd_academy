package handler

import (
	"banking_crowd/helper"
	"banking_crowd/models/classes"
	"banking_crowd/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type classHandler struct {
	classService service.ClassService
}

func NewClassHandler(classService service.ClassService) *classHandler {
	return &classHandler{classService}
}

func (h *classHandler) CreateClass(c *gin.Context) {
	var input classes.CreateClassInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create class", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newClass, err := h.classService.CreateClass(input)
	if err != nil {

		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success to create class", http.StatusOK, "success", classes.FormatInfoClass(newClass))
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
	response := helper.APIResponse("Success get all class", http.StatusOK, "success", classes.FormatInfoClasses(allclass))
	c.JSON(http.StatusOK, response)

}
