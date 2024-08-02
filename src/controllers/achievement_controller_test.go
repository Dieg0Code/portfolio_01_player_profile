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

func TestAchievementController_Create(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("CreateAchievement_Success", func(t *testing.T) {
		mockAchievementService := new(mocks.MockAchievementService)
		controller := NewAchievementController(mockAchievementService)
		router := gin.Default()
		router.POST("/achievement", controller.CreateAchievement)

		reqBody := request.CreateAchievementRequest{
			Name:        "Achievement 1",
			Description: "Description 1",
		}

		mockAchievementService.On("Create", reqBody).Return(nil)

		body, _ := json.Marshal(reqBody)
		req, err := http.NewRequest(http.MethodPost, "/achievement", bytes.NewBuffer(body))
		assert.NoError(t, err, "Expected no error creating request")
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code, "Status code should be 200")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Expected no error unmarshalling response")
		assert.Equal(t, 200, response.Code, "Response code should be 200")
		assert.Equal(t, "Success", response.Status, "Response status should be Success")

		mockAchievementService.AssertExpectations(t)
	})

	t.Run("CreateAchievement_InvalidRequestBody", func(t *testing.T) {
		mockAchievementService := new(mocks.MockAchievementService)
		controller := NewAchievementController(mockAchievementService)
		router := gin.Default()
		router.POST("/achievement", controller.CreateAchievement)

		req, err := http.NewRequest(http.MethodPost, "/achievement", bytes.NewBuffer([]byte("invalid json")))
		assert.NoError(t, err, "Expected no error creating request")
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code, "Status code should be 400")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Expected no error unmarshalling response")
		assert.Equal(t, 400, response.Code, "Response code should be 400")
		assert.Equal(t, "Error", response.Status, "Response status should be Error")

		mockAchievementService.AssertExpectations(t)
	})

	t.Run("CreateAchievement_FailedToCreateAchievement", func(t *testing.T) {
		mockAchievementService := new(mocks.MockAchievementService)
		controller := NewAchievementController(mockAchievementService)
		router := gin.Default()
		router.POST("/achievement", controller.CreateAchievement)

		reqBody := request.CreateAchievementRequest{
			Name:        "Achievement 1",
			Description: "Description 1",
		}

		mockAchievementService.On("Create", reqBody).Return(assert.AnError)

		body, _ := json.Marshal(reqBody)
		req, err := http.NewRequest(http.MethodPost, "/achievement", bytes.NewBuffer(body))
		assert.NoError(t, err, "Expected no error creating request")
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code, "Status code should be 500")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Expected no error unmarshalling response")
		assert.Equal(t, 500, response.Code, "Response code should be 500")
		assert.Equal(t, "Error", response.Status, "Response status should be Error")

		mockAchievementService.AssertExpectations(t)
	})
}

func TestAchievementController_GetAll(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("GetAllAchievements_Success", func(t *testing.T) {
		mockAchievementService := new(mocks.MockAchievementService)
		controller := NewAchievementController(mockAchievementService)
		router := gin.Default()
		router.GET("/achievement", controller.GetAllAchievements)

		mockAchievementService.On("GetAll", 1, 10).Return([]response.AchievementResponse{}, nil)

		req, err := http.NewRequest(http.MethodGet, "/achievement?page=1&pageSize=10", nil)
		assert.NoError(t, err, "Expected no error creating request")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code, "Status code should be 200")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Expected no error unmarshalling response")
		assert.Equal(t, 200, response.Code, "Response code should be 200")
		assert.Equal(t, "Success", response.Status, "Response status should be Success")

		mockAchievementService.AssertExpectations(t)
	})

	t.Run("GetAllAchievements_InvalidPage", func(t *testing.T) {
		mockAchievementService := new(mocks.MockAchievementService)
		controller := NewAchievementController(mockAchievementService)
		router := gin.Default()
		router.GET("/achievement", controller.GetAllAchievements)

		req, err := http.NewRequest(http.MethodGet, "/achievement?page=invalid&pageSize=10", nil)
		assert.NoError(t, err, "Expected no error creating request")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code, "Status code should be 400")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Expected no error unmarshalling response")
		assert.Equal(t, 400, response.Code, "Response code should be 400")
		assert.Equal(t, "Error", response.Status, "Response status should be Error")

		mockAchievementService.AssertExpectations(t)
	})

	t.Run("GetAllAchievements_InvalidPageSize", func(t *testing.T) {
		mockAchievementService := new(mocks.MockAchievementService)
		controller := NewAchievementController(mockAchievementService)
		router := gin.Default()
		router.GET("/achievement", controller.GetAllAchievements)

		req, err := http.NewRequest(http.MethodGet, "/achievement?page=1&pageSize=invalid", nil)
		assert.NoError(t, err, "Expected no error creating request")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code, "Status code should be 400")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Expected no error unmarshalling response")
		assert.Equal(t, 400, response.Code, "Response code should be 400")
		assert.Equal(t, "Error", response.Status, "Response status should be Error")

		mockAchievementService.AssertExpectations(t)

	})

	t.Run("GetAllAchievements_FailedToGetAllAchievements", func(t *testing.T) {
		mockAchievementService := new(mocks.MockAchievementService)
		controller := NewAchievementController(mockAchievementService)
		router := gin.Default()
		router.GET("/achievement", controller.GetAllAchievements)

		mockAchievementService.On("GetAll", 1, 10).Return([]response.AchievementResponse{}, assert.AnError)

		req, err := http.NewRequest(http.MethodGet, "/achievement?page=1&pageSize=10", nil)
		assert.NoError(t, err, "Expected no error creating request")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code, "Status code should be 500")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Expected no error unmarshalling response")
		assert.Equal(t, 500, response.Code, "Response code should be 500")
		assert.Equal(t, "Error", response.Status, "Response status should be Error")

		mockAchievementService.AssertExpectations(t)
	})
}

func TestAchievementController_GetAchivementByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("GetAchievementByID_Success", func(t *testing.T) {
		mockAchievementService := new(mocks.MockAchievementService)
		controller := NewAchievementController(mockAchievementService)
		router := gin.Default()
		router.GET("/achievement/:achievementID", controller.GetAchievementByID)

		mockAchievementService.On("GetByID", uint(1)).Return(&response.AchievementResponse{}, nil)

		req, err := http.NewRequest(http.MethodGet, "/achievement/1", nil)
		assert.NoError(t, err, "Expected no error creating request")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code, "Status code should be 200")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Expected no error unmarshalling response")
		assert.Equal(t, 200, response.Code, "Response code should be 200")
		assert.Equal(t, "Success", response.Status, "Response status should be Success")

		mockAchievementService.AssertExpectations(t)
	})

	t.Run("GetAchievementByID_InvalidAchievementID", func(t *testing.T) {
		mockAchievementService := new(mocks.MockAchievementService)
		controller := NewAchievementController(mockAchievementService)
		router := gin.Default()
		router.GET("/achievement/:achievementID", controller.GetAchievementByID)

		req, err := http.NewRequest(http.MethodGet, "/achievement/invalid", nil)
		assert.NoError(t, err, "Expected no error creating request")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code, "Status code should be 400")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Expected no error unmarshalling response")
		assert.Equal(t, 400, response.Code, "Response code should be 400")
		assert.Equal(t, "Error", response.Status, "Response status should be Error")

		mockAchievementService.AssertExpectations(t)
	})

	t.Run("GetAchievementByID_FailedToGetAchievementByID", func(t *testing.T) {
		mockAchievementService := new(mocks.MockAchievementService)
		controller := NewAchievementController(mockAchievementService)
		router := gin.Default()
		router.GET("/achievement/:achievementID", controller.GetAchievementByID)

		mockAchievementService.On("GetByID", uint(1)).Return(&response.AchievementResponse{}, assert.AnError)

		req, err := http.NewRequest(http.MethodGet, "/achievement/1", nil)
		assert.NoError(t, err, "Expected no error creating request")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code, "Status code should be 500")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Expected no error unmarshalling response")
		assert.Equal(t, 500, response.Code, "Response code should be 500")
		assert.Equal(t, "Error", response.Status, "Response status should be Error")

		mockAchievementService.AssertExpectations(t)
	})
}

func TestAchievementController_Update(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("UpdateAchievement_Success", func(t *testing.T) {
		mockAchievementService := new(mocks.MockAchievementService)
		controller := NewAchievementController(mockAchievementService)
		router := gin.Default()
		router.PUT("/achievement/:achievementID", controller.UpdateAchievement)

		reqBody := request.UpdateAchievementRequest{
			Name:        "Achievement 1",
			Description: "Description 1",
		}

		mockAchievementService.On("Update", uint(1), reqBody).Return(nil)

		body, _ := json.Marshal(reqBody)
		req, err := http.NewRequest(http.MethodPut, "/achievement/1", bytes.NewBuffer(body))
		assert.NoError(t, err, "Expected no error creating request")
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code, "Status code should be 200")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Expected no error unmarshalling response")
		assert.Equal(t, 200, response.Code, "Response code should be 200")
		assert.Equal(t, "Success", response.Status, "Response status should be Success")

		mockAchievementService.AssertExpectations(t)
	})

	t.Run("UpdateAchievement_InvalidRequestBody", func(t *testing.T) {
		mockAchievementService := new(mocks.MockAchievementService)
		controller := NewAchievementController(mockAchievementService)
		router := gin.Default()
		router.PUT("/achievement/:achievementID", controller.UpdateAchievement)

		req, err := http.NewRequest(http.MethodPut, "/achievement/1", bytes.NewBuffer([]byte("invalid json")))
		assert.NoError(t, err, "Expected no error creating request")
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code, "Status code should be 400")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Expected no error unmarshalling response")
		assert.Equal(t, 400, response.Code, "Response code should be 400")
		assert.Equal(t, "Error", response.Status, "Response status should be Error")

		mockAchievementService.AssertExpectations(t)
	})

	t.Run("UpdateAchievement_InvalidAchievementID", func(t *testing.T) {
		mockAchievementService := new(mocks.MockAchievementService)
		controller := NewAchievementController(mockAchievementService)
		router := gin.Default()
		router.PUT("/achievement/:achievementID", controller.UpdateAchievement)

		reqBody := request.UpdateAchievementRequest{
			Name:        "Achievement 1",
			Description: "Description 1",
		}

		body, _ := json.Marshal(reqBody)
		req, err := http.NewRequest(http.MethodPut, "/achievement/invalid", bytes.NewBuffer(body))
		assert.NoError(t, err, "Expected no error creating request")
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code, "Status code should be 400")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Expected no error unmarshalling response")
		assert.Equal(t, 400, response.Code, "Response code should be 400")
		assert.Equal(t, "Error", response.Status, "Response status should be Error")

		mockAchievementService.AssertExpectations(t)
	})

	t.Run("UpdateAchievement_FailedToUpdateAchievement", func(t *testing.T) {
		mockAchievementService := new(mocks.MockAchievementService)
		controller := NewAchievementController(mockAchievementService)
		router := gin.Default()
		router.PUT("/achievement/:achievementID", controller.UpdateAchievement)

		reqBody := request.UpdateAchievementRequest{
			Name:        "Achievement 1",
			Description: "Description 1",
		}

		mockAchievementService.On("Update", uint(1), reqBody).Return(assert.AnError)

		body, _ := json.Marshal(reqBody)
		req, err := http.NewRequest(http.MethodPut, "/achievement/1", bytes.NewBuffer(body))
		assert.NoError(t, err, "Expected no error creating request")
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code, "Status code should be 500")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Expected no error unmarshalling response")
		assert.Equal(t, 500, response.Code, "Response code should be 500")
		assert.Equal(t, "Error", response.Status, "Response status should be Error")

		mockAchievementService.AssertExpectations(t)
	})
}

func TestAchievementController_Delete(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("DeleteAchievement_Success", func(t *testing.T) {
		mockAchievementService := new(mocks.MockAchievementService)
		controller := NewAchievementController(mockAchievementService)
		router := gin.Default()
		router.DELETE("/achievement/:achievementID", controller.DeleteAchievement)

		mockAchievementService.On("Delete", uint(1)).Return(nil)

		req, err := http.NewRequest(http.MethodDelete, "/achievement/1", nil)
		assert.NoError(t, err, "Expected no error creating request")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code, "Status code should be 200")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Expected no error unmarshalling response")
		assert.Equal(t, 200, response.Code, "Response code should be 200")
		assert.Equal(t, "Success", response.Status, "Response status should be Success")

		mockAchievementService.AssertExpectations(t)
	})

	t.Run("DeleteAchievement_InvalidAchievementID", func(t *testing.T) {
		mockAchievementService := new(mocks.MockAchievementService)
		controller := NewAchievementController(mockAchievementService)
		router := gin.Default()
		router.DELETE("/achievement/:achievementID", controller.DeleteAchievement)

		req, err := http.NewRequest(http.MethodDelete, "/achievement/invalid", nil)
		assert.NoError(t, err, "Expected no error creating request")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code, "Status code should be 400")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Expected no error unmarshalling response")
		assert.Equal(t, 400, response.Code, "Response code should be 400")
		assert.Equal(t, "Error", response.Status, "Response status should be Error")

		mockAchievementService.AssertExpectations(t)
	})

	t.Run("DeleteAchievement_FailedToDeleteAchievement", func(t *testing.T) {
		mockAchievementService := new(mocks.MockAchievementService)
		controller := NewAchievementController(mockAchievementService)
		router := gin.Default()
		router.DELETE("/achievement/:achievementID", controller.DeleteAchievement)

		mockAchievementService.On("Delete", uint(1)).Return(assert.AnError)

		req, err := http.NewRequest(http.MethodDelete, "/achievement/1", nil)
		assert.NoError(t, err, "Expected no error creating request")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code, "Status code should be 500")

		var response response.BaseResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Expected no error unmarshalling response")
		assert.Equal(t, 500, response.Code, "Response code should be 500")
		assert.Equal(t, "Error", response.Status, "Response status should be Error")

		mockAchievementService.AssertExpectations(t)
	})
}

func TestAchievementController_GetAchievementWithPlayers(t *testing.T) {
	t.Run("GetAchievementWithPlayers_Success", func(t *testing.T) {
		mockAchievementService := new(mocks.MockAchievementService)
		controller := NewAchievementController(mockAchievementService)
		router := gin.Default()
		router.GET("/achievements/:achievementID/players", controller.GetAchievementWithPlayers)

		mockAchievementService.On("GetAchievementWithPlayers", uint(1)).Return(&response.AchievementWithPlayers{}, nil)

		req, _ := http.NewRequest(http.MethodGet, "/achievements/1/players", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code, "Status code should be 200")

		var response response.BaseResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Should be able to unmarshal response")
		assert.Equal(t, "Success", response.Status)

		mockAchievementService.AssertExpectations(t)
	})

	t.Run("GetAchievementWithPlayers_InvalidID", func(t *testing.T) {
		mockAchievementService := new(mocks.MockAchievementService)
		controller := NewAchievementController(mockAchievementService)
		router := gin.Default()
		router.GET("/achievements/:achievementID/players", controller.GetAchievementWithPlayers)

		req, _ := http.NewRequest(http.MethodGet, "/achievements/asd/players", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code, "Status code should be 400")
		mockAchievementService.AssertExpectations(t)
	})

	t.Run("GetAchievementWithPlayers_FailedToGetAchievementWithPlayers", func(t *testing.T) {
		mockAchievementService := new(mocks.MockAchievementService)
		controller := NewAchievementController(mockAchievementService)
		router := gin.Default()
		router.GET("/achievements/:achievementID/players", controller.GetAchievementWithPlayers)

		mockAchievementService.On("GetAchievementWithPlayers", uint(1)).Return(nil, assert.AnError)

		req, _ := http.NewRequest(http.MethodGet, "/achievements/1/players", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code, "Status code should be 500")
		mockAchievementService.AssertExpectations(t)
	})
}
