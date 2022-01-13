package test

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-playground/assert/v2"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
	"net/http/httptest"
	"pengdst/golang-restful-api/app"
	"pengdst/golang-restful-api/controller"
	"pengdst/golang-restful-api/helper"
	"pengdst/golang-restful-api/middleware"
	"pengdst/golang-restful-api/repository"
	"pengdst/golang-restful-api/service"
	"strings"
	"testing"
	"time"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang-restful-api-test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()

	categortRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categortRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func truncateDb(db *sql.DB) {
	db.Exec("TRUNCATE category")
}

func TestCreateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateDb(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name":"New Gadget"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	fmt.Println(body)
	var responseBody map[string]interface{}
	_ = json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])
	assert.NotEqual(t, nil, responseBody["data"])
}

func TestCreateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name":""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	_ = json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad Request", responseBody["status"])
	fmt.Println(responseBody)
}

func TestGetAllSuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	TestCreateCategorySuccess(t)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/categories", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	_ = json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])

	fmt.Println(responseBody)
	categories := responseBody["data"].([]interface{})
	assert.NotEqual(t, nil, categories[0].(map[string]interface{})["id"])
	assert.NotEqual(t, nil, categories[0].(map[string]interface{})["name"])
}

func TestFindByIdSuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	TestCreateCategorySuccess(t)

	requestBody := strings.NewReader(``)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/categories/1", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	_ = json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])
	assert.NotEqual(t, nil, responseBody["data"])
}

func TestFindByIdFailed(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	TestCreateCategorySuccess(t)

	requestBody := strings.NewReader(``)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/categories/100", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusNotFound, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	_ = json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not Found", responseBody["status"])
}

func TestUpdateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	TestCreateCategorySuccess(t)

	requestBody := strings.NewReader(`{"name":"Old Gadget"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/categories/1", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	_ = json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])
	assert.NotEqual(t, nil, responseBody["data"])
}

func TestUpdateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	TestCreateCategorySuccess(t)

	requestBody := strings.NewReader(`{"name":""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/categories/1", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	_ = json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad Request", responseBody["status"])
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	TestCreateCategorySuccess(t)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/categories/1", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	_ = json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])
}

func TestDeleteCategoryFailed(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	TestCreateCategorySuccess(t)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/categories/100", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusNotFound, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	_ = json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)

	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not Found", responseBody["status"])
}

func TestAuthorized(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	TestCreateCategorySuccess(t)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/categories", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	_ = json.Unmarshal(body, &responseBody)

	fmt.Println(response)
	assert.Equal(t, http.StatusUnauthorized, int(responseBody["code"].(float64)))
}
