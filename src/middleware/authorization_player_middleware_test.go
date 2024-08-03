package middleware

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dieg0code/player-profile/src/data/response"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRoleCheckPlayersMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Admin bypasses checks", func(t *testing.T) {
		router := setupRouter("admin", 1)

		mockGetPlayer := func(uint) (*response.PlayerProfileResponse, error) {
			return nil, nil // Esta función no debería ser llamada para admin
		}

		router.Use(RoleCheckPlayersMiddleware(mockGetPlayer))
		router.GET("/player/:playerID", func(c *gin.Context) {
			c.String(200, "OK")
		})

		w := performRequest(router, "GET", "/player/2")

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "OK")
	})

	t.Run("Valid player access", func(t *testing.T) {
		router := setupRouter("user", 1)

		mockGetPlayer := func(id uint) (*response.PlayerProfileResponse, error) {
			return &response.PlayerProfileResponse{UserID: 1}, nil
		}

		router.Use(RoleCheckPlayersMiddleware(mockGetPlayer))
		router.GET("/player/:playerID", func(c *gin.Context) {
			c.String(200, "OK")
		})

		w := performRequest(router, "GET", "/player/1")

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "OK")
	})

	t.Run("Invalid player access", func(t *testing.T) {
		router := setupRouter("user", 1)

		mockGetPlayer := func(id uint) (*response.PlayerProfileResponse, error) {
			return &response.PlayerProfileResponse{UserID: 2}, nil
		}

		router.Use(RoleCheckPlayersMiddleware(mockGetPlayer))
		router.GET("/player/:playerID", func(c *gin.Context) {
			c.String(200, "OK")
		})

		w := performRequest(router, "GET", "/player/2")

		assert.Equal(t, http.StatusForbidden, w.Code)
		assert.Contains(t, w.Body.String(), "You are not allowed to perform this action")
	})

	t.Run("Invalid player ID", func(t *testing.T) {
		router := setupRouter("user", 1)

		mockGetPlayer := func(uint) (*response.PlayerProfileResponse, error) {
			return nil, nil
		}

		router.Use(RoleCheckPlayersMiddleware(mockGetPlayer))
		router.GET("/player/:playerID", func(c *gin.Context) {
			c.String(200, "OK")
		})

		w := performRequest(router, "GET", "/player/invalid")

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Invalid player ID")
	})

	t.Run("Player not found", func(t *testing.T) {
		router := setupRouter("user", 1)

		mockGetPlayer := func(uint) (*response.PlayerProfileResponse, error) {
			return nil, errors.New("Player not found")
		}

		router.Use(RoleCheckPlayersMiddleware(mockGetPlayer))
		router.GET("/player/:playerID", func(c *gin.Context) {
			c.String(200, "OK")
		})

		w := performRequest(router, "GET", "/player/1")

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Contains(t, w.Body.String(), "Player not found")
	})
}

func setupRouter(role string, userID uint) *gin.Engine {
	router := gin.New()
	router.Use(func(c *gin.Context) {
		c.Set("role", role)
		c.Set("userID", userID)
	})
	return router
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
