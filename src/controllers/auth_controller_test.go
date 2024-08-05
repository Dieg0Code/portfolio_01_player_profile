package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dieg0code/player-profile/src/data/request"
	"github.com/dieg0code/player-profile/src/data/response"
	"github.com/dieg0code/player-profile/src/testutils/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthController_Login(t *testing.T) {
	mockAuthService := new(mocks.MockAuthService)
	authController := NewAuthController(mockAuthService)

	router := gin.Default()
	router.POST("/login", authController.Login)

	t.Run("Login_EmptyRequest", func(t *testing.T) {
		invalidReq := request.LoginRequest{
			Email:    "",
			Password: "",
		}

		body, err := json.Marshal(invalidReq)
		assert.Nil(t, err)
		req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
		assert.NoError(t, err, "Expected no error creating request")
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code, "Expected status code 400")
	})

	t.Run("Login_Success", func(t *testing.T) {
		loginReq := request.LoginRequest{
			Email:    "test@test.com",
			Password: "password123456",
		}
		mockAuthService.On("Login", loginReq).Return(&response.LoginResponse{
			Token: "token",
		}, nil)

		body, err := json.Marshal(loginReq)
		assert.Nil(t, err)
		req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
		assert.NoError(t, err, "Expected no error creating request")

		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code, "Expected status code 200")

		mockAuthService.AssertExpectations(t)
	})
}
