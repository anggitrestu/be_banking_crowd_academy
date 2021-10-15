package handler

import (
	"banking_crowd/helper"
	"banking_crowd/models/articles"
	"banking_crowd/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type articleHandler struct {
	articleService service.ArticleService
}

func NewArticleHandler(articleService service.ArticleService) *articleHandler {
	return &articleHandler{articleService}
}

func (h *articleHandler) CreateArticle(c *gin.Context) {
	var input articles.CreateArticleInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create class", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newArticle, err := h.articleService.CreateArticle(input)
	if err != nil {

		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success to create class", http.StatusOK, "success", articles.FormatArticle(newArticle))
	c.JSON(http.StatusOK, response)
}

func (h *articleHandler) GetAll(c *gin.Context) {
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
	allclass, err := h.articleService.GetAll(id)
	response := helper.APIResponse("Success get all class", http.StatusOK, "success", articles.FormatArticles(allclass))
	c.JSON(http.StatusOK, response)

}

func (h *articleHandler) Delete(c *gin.Context) {
	var inputID articles.GetArticleInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed delete article", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	errs := h.articleService.Delete(inputID.ID)
	if errs != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update learner", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	response := helper.APIResponse("Success delete article", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)

}
