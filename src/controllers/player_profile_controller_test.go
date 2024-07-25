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

func TestPlayerController_Create(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("CreatePlayer_Success", func(t *testing.T) {
		mockPlayerService := new(mocks.MockPlayerProfileService)
		controller := NewPlayerProfileController(mockPlayerService)
		router := gin.Default()
		router.POST("/player", controller.CreatePlayerProfile)

		reqBody := request.CreatePlayerProfileRequest{
			Nickname:   "dieg0",
			Avatar:     "https://avatar.com",
			Level:      1,
			Experience: 10,
			Points:     100,
			UserID:     1,
		}

		mockPlayerService.On("Create", reqBody).Return(nil)

		body, _ := json.Marshal(reqBody)
		req, err := http.NewRequest(http.MethodPost, "/player", bytes.NewBuffer(body))
		assert.NoError(t, err, "Expected no error in creating request")
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code, "Status code should be 200")
		mockPlayerService.AssertExpectations(t)
	})

	t.Run("CreatePlayer_InvalidRequestBody", func(t *testing.T) {
		mockPlayerService := new(mocks.MockPlayerProfileService)
		controller := NewPlayerProfileController(mockPlayerService)
		router := gin.Default()
		router.POST("/player", controller.CreatePlayerProfile)

		req, err := http.NewRequest(http.MethodPost, "/player", bytes.NewBuffer([]byte("invalid json")))
		assert.NoError(t, err, "Expected no error in creating request")
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code, "Status code should be 400")
		mockPlayerService.AssertNotCalled(t, "Create")
	})

	t.Run("CreatePlayer_FailedToCreate", func(t *testing.T) {
		mockPlayerService := new(mocks.MockPlayerProfileService)
		controller := NewPlayerProfileController(mockPlayerService)
		router := gin.Default()
		router.POST("/player", controller.CreatePlayerProfile)

		reqBody := request.CreatePlayerProfileRequest{
			Nickname:   "",
			Avatar:     "https://avatar.com",
			Level:      0,
			Experience: 10,
			Points:     100,
			UserID:     1,
		}

		mockPlayerService.On("Create", reqBody).Return(assert.AnError)

		body, _ := json.Marshal(reqBody)
		req, err := http.NewRequest(http.MethodPost, "/player", bytes.NewBuffer(body))
		assert.NoError(t, err, "Expected no error in creating request")
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code, "Status code should be 500")
		mockPlayerService.AssertExpectations(t)
	})
}

func TestPlayerController_GetAll(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("GetAllPlayers_Success", func(t *testing.T) {
		mockPlayerService := new(mocks.MockPlayerProfileService)
		controller := NewPlayerProfileController(mockPlayerService)
		router := gin.Default()
		router.GET("/players", controller.GetAllPlayers)

		mockPlayerService.On("GetAll", 1, 10).Return([]response.PlayerProfileResponse{}, nil)

		req, _ := http.NewRequest(http.MethodGet, "/players?page=1&pageSize=10", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code, "Status code should be 200")

		var response response.BaseResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Should be able to unmarshal response")
		assert.Equal(t, "Success", response.Status)

		mockPlayerService.AssertExpectations(t)
	})

	t.Run("GetAllPlayers_InvalidQueryParams", func(t *testing.T) {
		mockPlayerService := new(mocks.MockPlayerProfileService)
		controller := NewPlayerProfileController(mockPlayerService)
		router := gin.Default()
		router.GET("/players", controller.GetAllPlayers)

		req, _ := http.NewRequest(http.MethodGet, "/players?page=asd&pageSize=10", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code, "Status code should be 400")
		mockPlayerService.AssertNotCalled(t, "GetAll")
	})

	t.Run("GetAllPlayers_FailedToGetAll", func(t *testing.T) {
		mockPlayerService := new(mocks.MockPlayerProfileService)
		controller := NewPlayerProfileController(mockPlayerService)
		router := gin.Default()
		router.GET("/players", controller.GetAllPlayers)

		mockPlayerService.On("GetAll", 1, 10).Return(nil, assert.AnError)

		req, _ := http.NewRequest(http.MethodGet, "/players?page=1&pageSize=10", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code, "Status code should be 500")
		mockPlayerService.AssertExpectations(t)
	})
}

func TestPlayerController_GetPlayerByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("GetPlayerByID_Success", func(t *testing.T) {
		mockPlayerService := new(mocks.MockPlayerProfileService)
		controller := NewPlayerProfileController(mockPlayerService)
		router := gin.Default()
		router.GET("/player/:playerID", controller.GetPlayerByID)

		mockPlayerService.On("GetByID", uint(1)).Return(&response.PlayerProfileResponse{}, nil)

		req, _ := http.NewRequest(http.MethodGet, "/player/1", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code, "Status code should be 200")

		var response response.BaseResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Should be able to unmarshal response")
		assert.Equal(t, "Success", response.Status)

		mockPlayerService.AssertExpectations(t)
	})

	t.Run("GetPlayerByID_InvalidPlayerID", func(t *testing.T) {
		mockPlayerService := new(mocks.MockPlayerProfileService)
		controller := NewPlayerProfileController(mockPlayerService)
		router := gin.Default()
		router.GET("/player/:playerID", controller.GetPlayerByID)

		req, _ := http.NewRequest(http.MethodGet, "/player/asd", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code, "Status code should be 400")
		mockPlayerService.AssertNotCalled(t, "GetByID")
	})

	t.Run("GetPlayerByID_FailedToGetByID", func(t *testing.T) {
		mockPlayerService := new(mocks.MockPlayerProfileService)
		controller := NewPlayerProfileController(mockPlayerService)
		router := gin.Default()
		router.GET("/player/:playerID", controller.GetPlayerByID)

		mockPlayerService.On("GetByID", uint(1)).Return(nil, assert.AnError)

		req, _ := http.NewRequest(http.MethodGet, "/player/1", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code, "Status code should be 500")
		mockPlayerService.AssertExpectations(t)
	})
}

func TestPlayerController_Update(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("UpdatePlayer_Success", func(t *testing.T) {
		mockPlayerService := new(mocks.MockPlayerProfileService)
		controller := NewPlayerProfileController(mockPlayerService)
		router := gin.Default()
		router.PUT("/player/:playerID", controller.UpdatePlayer)

		reqBody := request.UpdatePlayerProfileRequest{
			Nickname:   "dieg0",
			Avatar:     "https://avatar.com",
			Level:      1,
			Experience: 10,
			Points:     100,
		}

		mockPlayerService.On("Update", uint(1), reqBody).Return(nil)

		body, _ := json.Marshal(reqBody)
		req, err := http.NewRequest(http.MethodPut, "/player/1", bytes.NewBuffer(body))
		assert.NoError(t, err, "Expected no error in creating request")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code, "Status code should be 200")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Should be able to unmarshal response")
		assert.Equal(t, "Success", response.Status)

		mockPlayerService.AssertExpectations(t)
	})

	t.Run("UpdatePlayer_InvalidRequestBody", func(t *testing.T) {
		mockPlayerService := new(mocks.MockPlayerProfileService)
		controller := NewPlayerProfileController(mockPlayerService)
		router := gin.Default()
		router.PUT("/player/:playerID", controller.UpdatePlayer)

		req, err := http.NewRequest(http.MethodPut, "/player/1", bytes.NewBuffer([]byte("invalid json")))
		assert.NoError(t, err, "Expected no error in creating request")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code, "Status code should be 400")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Should be able to unmarshal response")
		assert.Equal(t, "Error", response.Status)

		mockPlayerService.AssertNotCalled(t, "Update")
	})

	t.Run("UpdatePlayer_InvalidPlayerID", func(t *testing.T) {
		mockPlayerService := new(mocks.MockPlayerProfileService)
		controller := NewPlayerProfileController(mockPlayerService)
		router := gin.Default()
		router.PUT("/player/:playerID", controller.UpdatePlayer)

		reqBody := request.UpdatePlayerProfileRequest{
			Nickname:   "dieg0",
			Avatar:     "https://avatar.com",
			Level:      1,
			Experience: 10,
			Points:     100,
		}

		body, _ := json.Marshal(reqBody)
		req, err := http.NewRequest(http.MethodPut, "/player/asd", bytes.NewBuffer(body))
		assert.NoError(t, err, "Expected no error in creating request")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code, "Status code should be 400")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Should be able to unmarshal response")
		assert.Equal(t, "Error", response.Status)

		mockPlayerService.AssertNotCalled(t, "Update")
	})
}

func TestPlayerController_Delete(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("DeletePlayer_Success", func(t *testing.T) {
		mockPlayerService := new(mocks.MockPlayerProfileService)
		controller := NewPlayerProfileController(mockPlayerService)
		router := gin.Default()
		router.DELETE("/player/:playerID", controller.DeletePlayer)

		mockPlayerService.On("Delete", uint(1)).Return(nil)

		req, _ := http.NewRequest(http.MethodDelete, "/player/1", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code, "Status code should be 200")

		var response response.BaseResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Should be able to unmarshal response")
		assert.Equal(t, "Success", response.Status)

		mockPlayerService.AssertExpectations(t)
	})

	t.Run("DeletePlayer_InvalidPlayerID", func(t *testing.T) {
		mockPlayerService := new(mocks.MockPlayerProfileService)
		controller := NewPlayerProfileController(mockPlayerService)
		router := gin.Default()
		router.DELETE("/player/:playerID", controller.DeletePlayer)

		req, _ := http.NewRequest(http.MethodDelete, "/player/asd", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code, "Status code should be 400")
		mockPlayerService.AssertNotCalled(t, "Delete")
	})

	t.Run("DeletePlayer_FailedToDelete", func(t *testing.T) {
		mockPlayerService := new(mocks.MockPlayerProfileService)
		controller := NewPlayerProfileController(mockPlayerService)
		router := gin.Default()
		router.DELETE("/player/:playerID", controller.DeletePlayer)

		mockPlayerService.On("Delete", uint(1)).Return(assert.AnError)

		req, _ := http.NewRequest(http.MethodDelete, "/player/1", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code, "Status code should be 500")

		var response response.BaseResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Should be able to unmarshal response")
		assert.Equal(t, "Error", response.Status)

		mockPlayerService.AssertExpectations(t)
	})
}
