package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dieg0code/player-profile/src/data/request"
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
