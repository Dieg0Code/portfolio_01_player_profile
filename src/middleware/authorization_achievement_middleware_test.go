package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthorizationAchievementsMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("Valid admin access", func(t *testing.T) {
		router := gin.New()
		router.Use(func(ctx *gin.Context) {
			ctx.Set("role", "admin")
			ctx.Next()
		})

		router.Use(AuthorizationAchievementMiddleware())
		router.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "success"})
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code, "response status should be 200")
		assert.Contains(t, w.Body.String(), "success", "response body should contain 'success'")
	})

	t.Run("Invalid user access", func(t *testing.T) {
		router := gin.New()
		router.Use(func(ctx *gin.Context) {
			ctx.Set("role", "user")
			ctx.Next()
		})

		router.Use(AuthorizationAchievementMiddleware())
		router.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "success"})
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, 403, w.Code, "response status should be 403")
		assert.Contains(t, w.Body.String(), "You are not allowed to perform this action", "response body should contain 'You are not allowed to perform this action'")
	})
}
