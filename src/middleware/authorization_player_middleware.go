package middleware

import (
	"strconv"

	"github.com/dieg0code/player-profile/src/data/response"
	"github.com/gin-gonic/gin"
)

type GetPlayerFunc func(uint) (*response.PlayerProfileResponse, error)

func RoleCheckPlayersMiddleware(getPlayer GetPlayerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authUserID := ctx.GetUint("userID")
		authUserRole := ctx.GetString("role")

		if authUserRole == "admin" {
			ctx.Next()
			return
		}

		playerID := ctx.Param("playerID")
		playerIDUint, err := strconv.ParseUint(playerID, 10, 64)
		if err != nil {
			ctx.JSON(400, response.BaseResponse{
				Code:    400,
				Status:  "Bad Request",
				Message: "Invalid player ID",
				Data:    nil,
			})
			ctx.Abort()
			return
		}

		player, err := getPlayer(uint(playerIDUint))
		if err != nil {
			ctx.JSON(404, response.BaseResponse{
				Code:    404,
				Status:  "Not Found",
				Message: "Player not found",
				Data:    nil,
			})
			ctx.Abort()
			return
		}

		if player.UserID != authUserID {
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
