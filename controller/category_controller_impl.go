package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pengdst/golang-restful-api/helper"
	"pengdst/golang-restful-api/model/web"
	"pengdst/golang-restful-api/service"
	"strconv"
)

type CategoryControllerImpl struct {
	Service service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		Service: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(c *gin.Context) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	c.ShouldBind(&categoryCreateRequest)

	categoryResponse := controller.Service.Create(c, categoryCreateRequest)
	response := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "Ok",
		Data:   categoryResponse,
	}

	c.JSON(http.StatusCreated, response)
}

func (controller *CategoryControllerImpl) Update(c *gin.Context) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	c.ShouldBind(&categoryUpdateRequest)

	categoryId := c.Param("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	categoryResponse := controller.Service.Update(c, categoryUpdateRequest)
	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   categoryResponse,
	}

	c.JSON(http.StatusOK, response)
}

func (controller *CategoryControllerImpl) Delete(c *gin.Context) {
	categoryId := c.Param("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.Service.Delete(c, id)
	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
	}

	c.JSON(http.StatusOK, response)
}

func (controller *CategoryControllerImpl) FindById(c *gin.Context) {
	categoryId := c.Param("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := controller.Service.FindById(c, id)
	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   categoryResponse,
	}

	c.JSON(http.StatusOK, response)
}

func (controller *CategoryControllerImpl) GetAll(c *gin.Context) {
	categoryResponses := controller.Service.GetAll(c)
	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   categoryResponses,
	}

	c.JSON(http.StatusOK, response)
}
