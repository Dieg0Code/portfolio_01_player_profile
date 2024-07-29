package middleware

import (
	"strconv"

	"github.com/dieg0code/player-profile/src/data/response"
	"github.com/gin-gonic/gin"
)

// RoleCheckMiddleware verifies user roles and permissions.
func RoleCheckUsersMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authUserID := ctx.GetUint("userID")
		authUserRole := ctx.GetString("role")

		if authUserRole == "admin" {
			ctx.Next()
			return
		}

		userID := ctx.Param("userID")
		userIDUint, err := strconv.ParseUint(userID, 10, 64)
		if err != nil {
			ctx.JSON(400, response.BaseResponse{
				Code:    400,
				Status:  "Bad Request",
				Message: "Invalid user ID",
				Data:    nil,
			})
			ctx.Abort()
			return
		}

		if authUserID != uint(userIDUint) {
			ctx.JSON(403, response.BaseResponse{
				Code:    403,
				Status:  "Forbidden",
				Message: "You are not allowed to perform this action",
				Data:    nil,
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
