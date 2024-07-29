package middleware

import (
	"strings"

	"github.com/dieg0code/player-profile/src/auth"
	"github.com/dieg0code/player-profile/src/data/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" {
			errorResponse := response.BaseResponse{
				Code:    401,
				Status:  "Unauthorized",
				Message: "Token is required",
				Data:    nil,
			}

			ctx.JSON(401, errorResponse)
			ctx.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := auth.ParseToken(tokenString)
		if err != nil || !token.Valid {
			errorResponse := response.BaseResponse{
				Code:    401,
				Status:  "Unauthorized",
				Message: "Invalid token",
				Data:    nil,
			}

			ctx.JSON(401, errorResponse)
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			errorResponse := response.BaseResponse{
				Code:    401,
				Status:  "Unauthorized",
				Message: "Invalid token",
				Data:    nil,
			}

			ctx.JSON(401, errorResponse)
			ctx.Abort()
			return
		}

		userID := uint(claims["userID"].(float64))
		role, ok := claims["role"].(string)
		if !ok {
			errorResponse := response.BaseResponse{
				Code:    401,
				Status:  "Unauthorized",
				Message: "Invalid token",
				Data:    nil,
			}

			ctx.JSON(401, errorResponse)
			ctx.Abort()
			return
		}
		ctx.Set("userID", userID)
		ctx.Set("role", role)

		ctx.Next()
	}
}
