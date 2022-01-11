package exception

import (
	"net/http"
	"pengdst/golang-restful-api/helper"
	"pengdst/golang-restful-api/model/web"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, error interface{}) {

	if notFoundError(writer, request, error) {
		return
	}

	if validationError(writer, request, error) {
		return
	}

	internalServerError(writer, request, error)
}

func validationError(writer http.ResponseWriter, request *http.Request, error interface{}) bool {
	exception, ok := error.(ValidationError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
	}

	return ok
}

func notFoundError(writer http.ResponseWriter, request *http.Request, error interface{}) bool {
	exception, ok := error.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
	}

	return ok
}

func internalServerError(writer http.ResponseWriter, request *http.Request, error interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   error,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
