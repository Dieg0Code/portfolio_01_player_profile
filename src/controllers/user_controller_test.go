package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dieg0code/player-profile/src/data/request"
	"github.com/dieg0code/player-profile/src/data/response"
	"github.com/dieg0code/player-profile/src/testutils/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserController_Create(t *testing.T) {
	mockUserService := new(mocks.MockUserService)
	controller := NewUserController(mockUserService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/users", controller.CreateUser)

	t.Run("CreaterUser_Success", func(t *testing.T) {
		reqBody := request.CreateUserRequest{
			UserName: "test",
			Email:    "test@test.com",
			Password: "password",
			Age:      20,
		}

		mockUserService.On("Create", reqBody).Return(nil)

		body, _ := json.Marshal(reqBody)
		req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
		assert.NoError(t, err, "Expected no error in creating request")
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code, "Status code should be 200")
		mockUserService.AssertExpectations(t)

	})

	t.Run("CreateUser_InvalidRequestBody", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer([]byte("invalid json")))
		assert.NoError(t, err, "Expected no error in creating request")
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code, "Status code should be 400")
	})

	t.Run("CreateUser_FailedToCreateUser", func(t *testing.T) {
		reqBody := request.CreateUserRequest{
			UserName: "test",
			Email:    "test@test.com",
			Age:      20,
		}

		mockUserService.On("Create", reqBody).Return(assert.AnError)

		body, _ := json.Marshal(reqBody)
		req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
		assert.NoError(t, err, "Expected no error in creating request")
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code, "Expected 500 creating a invalid user")
		mockUserService.AssertExpectations(t)
	})
}

func TestUserController_GetAllUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("GetAllUsers_Success", func(t *testing.T) {
		mockUserService := new(mocks.MockUserService)
		controller := NewUserController(mockUserService)
		router := gin.Default()
		router.GET("/users", controller.GetAllUsers)

		mockUserService.On("GetAll", 1, 10).Return([]response.UserResponse{}, nil)

		req, _ := http.NewRequest(http.MethodGet, "/users?page=1&pageSize=10", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code, "Status code should be 200")

		var response response.BaseResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Should be able to unmarshal response")
		assert.Equal(t, "Success", response.Status)

		mockUserService.AssertExpectations(t)
	})

	t.Run("GetAllUsers_InvalidPagination", func(t *testing.T) {
		mockUserService := new(mocks.MockUserService)
		controller := NewUserController(mockUserService)
		router := gin.Default()
		router.GET("/users", controller.GetAllUsers)

		req, _ := http.NewRequest(http.MethodGet, "/users?page=asd&pageSize=asd", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code, "Status code should be 400")

		var response response.BaseResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Should be able to unmarshal response")
		assert.Equal(t, "Error", response.Status)
		assert.Equal(t, "Invalid page", response.Message)
	})

	t.Run("GetAllUsers_FailedToGetUsers", func(t *testing.T) {
		mockUserService := new(mocks.MockUserService)
		controller := NewUserController(mockUserService)
		router := gin.Default()
		router.GET("/users", controller.GetAllUsers)

		mockUserService.On("GetAll", 1, 10).Return([]response.UserResponse{}, errors.New("Service Error"))

		req, _ := http.NewRequest(http.MethodGet, "/users?page=1&pageSize=10", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code, "Status code should be 500")

		var response response.BaseResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Should be able to unmarshal response")
		assert.Equal(t, "Error", response.Status)
		assert.Equal(t, "Failed to get users", response.Message)

		mockUserService.AssertExpectations(t)
	})
}

func TestUserController_GetByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("GetByID_Success", func(t *testing.T) {
		mockUserService := new(mocks.MockUserService)
		controller := NewUserController(mockUserService)
		router := gin.Default()
		router.GET("/users/:userID", controller.GetUserByID)

		mockUserService.On("GetByID", uint(1)).Return(&response.UserResponse{}, nil)

		req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code, "Status code should be 200")

		var response response.BaseResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Should be able to unmarshal response")
		assert.Equal(t, "Success", response.Status)

		mockUserService.AssertExpectations(t)
	})

	t.Run("GetByID_InvalidUserID", func(t *testing.T) {
		mockUserService := new(mocks.MockUserService)
		controller := NewUserController(mockUserService)
		router := gin.Default()
		router.GET("/users/:userID", controller.GetUserByID)

		req, _ := http.NewRequest(http.MethodGet, "/users/asd", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code, "Status code should be 400")

		var response response.BaseResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Should be able to unmarshal response")
		assert.Equal(t, "Error", response.Status)
		assert.Equal(t, "invalid user id", response.Message)
	})

	t.Run("UserNotFound", func(t *testing.T) {
		mockUserService := new(mocks.MockUserService)
		controller := NewUserController(mockUserService)
		router := gin.Default()
		router.GET("/users/:userID", controller.GetUserByID)
		mockUserService.On("GetByID", uint(1)).Return(&response.UserResponse{}, errors.New("user not found"))

		req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)

		var response response.BaseResponse
		err := json.Unmarshal(resp.Body.Bytes(), &response)
		assert.NoError(t, err, "Should be able to unmarshal response")
		assert.Equal(t, "Error", response.Status)
		assert.Equal(t, "Failed to get user", response.Message)
		assert.Nil(t, response.Data)

		mockUserService.AssertExpectations(t)
	})
}
