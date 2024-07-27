package controllers

import (
	"strconv"

	"github.com/dieg0code/player-profile/src/data/request"
	"github.com/dieg0code/player-profile/src/data/response"
	"github.com/dieg0code/player-profile/src/helpers"
	"github.com/dieg0code/player-profile/src/services"
	"github.com/gin-gonic/gin"
)

type PlayerProfileController struct {
	playerProfileService services.PlayerProfileService
}

func NewPlayerProfileController(service services.PlayerProfileService) *PlayerProfileController {
	return &PlayerProfileController{
		playerProfileService: service,
	}
}

func (controller *PlayerProfileController) CreatePlayerProfile(ctx *gin.Context) {
	createPlayerProfileRequest := request.CreatePlayerProfileRequest{}

	err := ctx.ShouldBindJSON(&createPlayerProfileRequest)
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

	err = controller.playerProfileService.Create(createPlayerProfileRequest)
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    500,
			Status:  "Error",
			Message: "Failed to create player profile",
			Data:    nil,
		}

		ctx.JSON(500, errorResponse)
		return
	}

	webResponse := response.BaseResponse{
		Code:    200,
		Status:  "Success",
		Message: "Player profile created successfully",
		Data:    nil,
	}

	ctx.JSON(200, webResponse)
}

func (controller *PlayerProfileController) GetAllPlayers(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	pageSize := ctx.DefaultQuery("pageSize", "10")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    400,
			Status:  "Error",
			Message: "Invalid page",
			Data:    nil,
		}

		ctx.JSON(400, errorResponse)
		return
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    400,
			Status:  "Error",
			Message: "Invalid pageSize",
			Data:    nil,
		}

		ctx.JSON(400, errorResponse)
		return
	}

	players, err := controller.playerProfileService.GetAll(pageInt, pageSizeInt)
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    500,
			Status:  "Error",
			Message: "Failed to get players",
			Data:    nil,
		}

		ctx.JSON(500, errorResponse)
		return
	}

	webResponse := response.BaseResponse{
		Code:    200,
		Status:  "Success",
		Message: "Players fetched successfully",
		Data:    players,
	}

	ctx.JSON(200, webResponse)
}

func (controller *PlayerProfileController) GetPlayerByID(ctx *gin.Context) {
	playerID := ctx.Param("playerID")

	playerIDInt, err := strconv.Atoi(playerID)
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    400,
			Status:  "Error",
			Message: helpers.ErrInvalidPlayerProfileID.Error(),
			Data:    nil,
		}

		ctx.JSON(400, errorResponse)
		return
	}

	player, err := controller.playerProfileService.GetByID(uint(playerIDInt))
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    500,
			Status:  "Error",
			Message: "Failed to get player",
			Data:    nil,
		}

		ctx.JSON(500, errorResponse)
		return
	}

	webResponse := response.BaseResponse{
		Code:    200,
		Status:  "Success",
		Message: "Player fetched successfully",
		Data:    player,
	}

	ctx.JSON(200, webResponse)
}

func (controller *PlayerProfileController) UpdatePlayer(ctx *gin.Context) {
	playerID := ctx.Param("playerID")
	updatePlayerProfileRequest := request.UpdatePlayerProfileRequest{}

	err := ctx.ShouldBindJSON(&updatePlayerProfileRequest)
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

	playerIDInt, err := strconv.Atoi(playerID)
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    400,
			Status:  "Error",
			Message: helpers.ErrInvalidPlayerProfileID.Error(),
			Data:    nil,
		}

		ctx.JSON(400, errorResponse)
		return
	}

	err = controller.playerProfileService.Update(uint(playerIDInt), updatePlayerProfileRequest)
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    500,
			Status:  "Error",
			Message: "Failed to update player",
			Data:    nil,
		}

		ctx.JSON(500, errorResponse)
		return
	}

	webResponse := response.BaseResponse{
		Code:    200,
		Status:  "Success",
		Message: "Player updated successfully",
		Data:    nil,
	}

	ctx.JSON(200, webResponse)
}

func (controller *PlayerProfileController) DeletePlayer(ctx *gin.Context) {
	playerID := ctx.Param("playerID")

	playerIDInt, err := strconv.Atoi(playerID)
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    400,
			Status:  "Error",
			Message: helpers.ErrInvalidPlayerProfileID.Error(),
			Data:    nil,
		}

		ctx.JSON(400, errorResponse)
		return
	}

	err = controller.playerProfileService.Delete(uint(playerIDInt))
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    500,
			Status:  "Error",
			Message: "Failed to delete player",
			Data:    nil,
		}

		ctx.JSON(500, errorResponse)
		return
	}

	webResponse := response.BaseResponse{
		Code:    200,
		Status:  "Success",
		Message: "Player deleted successfully",
		Data:    nil,
	}

	ctx.JSON(200, webResponse)
}
