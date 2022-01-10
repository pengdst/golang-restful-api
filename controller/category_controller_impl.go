package controller

import (
	"github.com/julienschmidt/httprouter"
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

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, categoryCreateRequest)

	categoryResponse := controller.Service.Create(request.Context(), categoryCreateRequest)
	response := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, response)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, categoryUpdateRequest)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	categoryResponse := controller.Service.Update(request.Context(), categoryUpdateRequest)
	response := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, response)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.Service.Delete(request.Context(), id)
	response := web.WebResponse{
		Code:   200,
		Status: "Ok",
	}

	helper.WriteToResponseBody(writer, response)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := controller.Service.FindById(request.Context(), id)
	response := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, response)
}

func (controller *CategoryControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponses := controller.Service.GetAll(request.Context())
	response := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponses,
	}

	helper.WriteToResponseBody(writer, response)
}
