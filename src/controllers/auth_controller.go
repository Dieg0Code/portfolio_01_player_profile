package controllers

import (
	"github.com/dieg0code/player-profile/src/data/request"
	"github.com/dieg0code/player-profile/src/data/response"
	"github.com/dieg0code/player-profile/src/services"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(service services.AuthService) *AuthController {
	return &AuthController{
		authService: service,
	}
}

// Login godoc
//	@Summary		Login to the application
//	@Description	Login to the application with the input payload
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.LoginRequest	true	"Login Request"
//	@Success		200		{object}	response.BaseResponse
//	@Failure		400		{object}	response.BaseResponse
//	@Failure		500		{object}	response.BaseResponse
//	@Router			/login [post]
func (controller *AuthController) Login(ctx *gin.Context) {
	loginRequest := request.LoginRequest{}

	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    400,
			Status:  "Error",
			Message: "Invalid request body",
			Data:    nil,
		}

		ctx.JSON(400, errorResponse)
		return
	}

	loginResponse, err := controller.authService.Login(loginRequest)
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    500,
			Status:  "Error",
			Message: "Failed to login",
			Data:    nil,
		}

		ctx.JSON(500, errorResponse)
		return
	}

	ctx.Header("Authorization", "Bearer "+loginResponse.Token)

	webResponse := response.BaseResponse{
		Code:    200,
		Status:  "Success",
		Message: "Login successful",
		Data:    nil,
	}

	ctx.JSON(200, webResponse)
}
