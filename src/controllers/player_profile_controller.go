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

// CreatePlayerProfile godoc
//
//	@Summary		Create a new player profile
//	@Description	Create a new player profile with the input payload
//	@Tags			Player
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.CreatePlayerProfileRequest	true	"Create Player Profile Request"
//	@Success		200		{object}	response.BaseResponse
//	@Failure		400		{object}	response.BaseResponse
//	@Failure		500		{object}	response.BaseResponse
//	@Router			/players [post]
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

// GetAllPlayers godoc
//
//	@Summary		Get all players
//	@Description	Get all players with pagination, by default page is 1 and pageSize is 10
//	@Tags			Player
//	@Accept			json
//	@Produce		json
//	@Param			page		query		int	false	"Page number"
//	@Param			pageSize	query		int	false	"Page size"
//	@Success		200			{object}	response.BaseResponse
//	@Failure		400			{object}	response.BaseResponse
//	@Failure		500			{object}	response.BaseResponse
//	@Router			/players [get]
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

// GetPlayerByID godoc
//
//	@Summary		Get player by ID
//	@Description	Get player by ID
//	@Tags			Player
//	@Accept			json
//	@Produce		json
//	@Param			playerID	path		int	true	"Player ID"
//	@Success		200			{object}	response.BaseResponse
//	@Failure		400			{object}	response.BaseResponse
//	@Failure		500			{object}	response.BaseResponse
//	@Router			/players/{playerID} [get]
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

// UpdatePlayer godoc
//
//	@Summary		Update player by ID
//	@Description	Update player by ID
//	@Tags			Player
//	@Accept			json
//	@Produce		json
//	@Param			playerID	path		int									true	"Player ID"
//	@Param			request		body		request.UpdatePlayerProfileRequest	true	"Update Player Profile Request"
//	@Success		200			{object}	response.BaseResponse
//	@Failure		400			{object}	response.BaseResponse
//	@Failure		500			{object}	response.BaseResponse
//	@Router			/players/{playerID} [put]
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

// DeletePlayer godoc
//
//	@Summary		Delete player by ID
//	@Description	Delete player by ID
//	@Tags			Player
//	@Accept			json
//	@Produce		json
//	@Param			playerID	path		int	true	"Player ID"
//	@Success		200			{object}	response.BaseResponse
//	@Failure		400			{object}	response.BaseResponse
//	@Failure		500			{object}	response.BaseResponse
//	@Router			/players/{playerID} [delete]
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

// GetPlayerWithAchievements godoc
//
//	@Summary		Get player with achievements by ID
//	@Description	Get player with achievements by ID
//	@Tags			Player
//	@Accept			json
//	@Produce		json
//	@Param			playerID	path		int	true	"Player ID"
//	@Success		200			{object}	response.BaseResponse
//	@Failure		400			{object}	response.BaseResponse
//	@Failure		500			{object}	response.BaseResponse
//	@Router			/players/{playerID}/achievements [get]
func (controller *PlayerProfileController) GetPlayerWithAchievements(ctx *gin.Context) {
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

	player, err := controller.playerProfileService.GetPlayerWithAchievements(uint(playerIDInt))
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    500,
			Status:  "Error",
			Message: "Failed to get player with achievements",
			Data:    nil,
		}

		ctx.JSON(500, errorResponse)
		return
	}

	webResponse := response.BaseResponse{
		Code:    200,
		Status:  "Success",
		Message: "Player with achievements fetched successfully",
		Data:    player,
	}

	ctx.JSON(200, webResponse)
}

func (controller *PlayerProfileController) GetPlayerByIDFromService(playerID uint) (*response.PlayerProfileResponse, error) {
	return controller.playerProfileService.GetByID(playerID)
}
