package middleware

import (
	"github.com/dieg0code/player-profile/src/data/response"
	"github.com/gin-gonic/gin"
)

func AuthorizationAchievementMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authUserRole := ctx.GetString("role")

		if authUserRole == "admin" {
			ctx.Next()
			return
		}

		ctx.JSON(403, response.BaseResponse{
			Code:    403,
			Status:  "Forbidden",
			Message: "You are not allowed to perform this action",
			Data:    nil,
		})

		ctx.Abort()
	}
}
