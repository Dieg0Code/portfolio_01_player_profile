package middleware

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("ValidToken", func(t *testing.T) {
		router := gin.New()
		router.Use(JWTAuthMiddleware())
		router.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "success"})
		})

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userID": 1,
			"role":   "admin",
		})

		tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
		assert.Nil(t, err, "Expected no error signing token")

		req, err := http.NewRequest(http.MethodGet, "/test", nil)
		assert.Nil(t, err, "Expected no error creating request")

		req.Header.Set("Authorization", "Bearer "+tokenString)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code, "Expected status code 200")
		assert.Contains(t, rec.Body.String(), "success", "Expected response body to contain 'success'")
	})

	t.Run("InvalidToken", func(t *testing.T) {
		router := gin.New()
		router.Use(JWTAuthMiddleware())
		router.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "success"})
		})

		req, err := http.NewRequest(http.MethodGet, "/test", nil)
		assert.Nil(t, err, "Expected no error creating request")

		req.Header.Set("Authorization", "Bearer invalidtoken")
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code, "Expected status code 401")
		assert.Contains(t, rec.Body.String(), "Unauthorized", "Expected response body to contain 'Unauthorized'")
	})

	t.Run("NoToken", func(t *testing.T) {
		router := gin.New()
		router.Use(JWTAuthMiddleware())
		router.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "success"})
		})

		req, err := http.NewRequest(http.MethodGet, "/test", nil)
		assert.Nil(t, err, "Expected no error creating request")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code, "Expected status code 401")
		assert.Contains(t, rec.Body.String(), "Unauthorized", "Expected response body to contain 'Unauthorized'")
	})

	t.Run("Valid token invalid claims", func(t *testing.T) {
		router := gin.New()
		router.Use(JWTAuthMiddleware())
		router.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"invalidClaim": "value",
		})
		tokenString, _ := token.SignedString([]byte("secret"))

		req, err := http.NewRequest(http.MethodGet, "/test", nil)
		assert.Nil(t, err, "Expected no error creating request")

		req.Header.Set("Authorization", "Bearer "+tokenString)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code, "Expected status code 401")
		assert.Contains(t, rec.Body.String(), "Invalid token", "Expected response body to contain 'Invalid token'")
	})

	t.Run("Invalid Prefix", func(t *testing.T) {
		router := gin.New()
		router.Use(JWTAuthMiddleware())
		router.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})

		req, err := http.NewRequest(http.MethodGet, "/test", nil)
		assert.Nil(t, err, "Expected no error creating request")

		req.Header.Set("Authorization", "InvalidPrefix token")
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code, "Expected status code 401")
		assert.Contains(t, rec.Body.String(), "Invalid token", "Expected response body to contain 'Invalid token'")
	})

	t.Run("Invalid role claim", func(t *testing.T) {
		router := gin.New()
		router.Use(JWTAuthMiddleware())
		router.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userID": 1,
			"role":   "invalid",
		})
		tokenString, _ := token.SignedString([]byte("secret"))

		req, err := http.NewRequest(http.MethodGet, "/test", nil)
		assert.Nil(t, err, "Expected no error creating request")

		req.Header.Set("Authorization", "Bearer "+tokenString)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code, "Expected status code 401")
		assert.Contains(t, rec.Body.String(), "Unauthorized", "Expected response body to contain 'Unauthorized'")
	})
}
