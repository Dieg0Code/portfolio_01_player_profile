package controllers

import (
	"strconv"

	"github.com/dieg0code/player-profile/src/data/request"
	"github.com/dieg0code/player-profile/src/data/response"
	"github.com/dieg0code/player-profile/src/services"
	"github.com/gin-gonic/gin"
)

type AchievementController struct {
	achievementService services.AchievementService
}

func NewAchievementController(service services.AchievementService) *AchievementController {
	return &AchievementController{
		achievementService: service,
	}
}

// CreateAchievement godoc
//
//	@Summary		Create a new achievement
//	@Description	Create a new achievement with the input payload
//	@Tags			Achievement
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.CreateAchievementRequest	true	"Create Achievement Request"
//	@Success		200		{object}	response.BaseResponse
//	@Failure		400		{object}	response.BaseResponse
//	@Failure		500		{object}	response.BaseResponse
//	@Router			/achievements [post]
func (controller *AchievementController) CreateAchievement(ctx *gin.Context) {
	createAchievementRequest := request.CreateAchievementRequest{}

	err := ctx.ShouldBindJSON(&createAchievementRequest)
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

	err = controller.achievementService.Create(createAchievementRequest)
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    500,
			Status:  "Error",
			Message: "Failed to create achievement",
			Data:    nil,
		}

		ctx.JSON(500, errorResponse)
		return
	}

	webResponse := response.BaseResponse{
		Code:    200,
		Status:  "Success",
		Message: "Achievement created successfully",
		Data:    nil,
	}

	ctx.JSON(200, webResponse)
}

// GetAllAchievements godoc
//
//	@Summary		Get all achievements
//	@Description	Get all achievements with pagination, default page is 1 and default pageSize is 10
//	@Tags			Achievement
//	@Accept			json
//	@Produce		json
//	@Param			page		query		int	false	"Page number"
//	@Param			pageSize	query		int	false	"Page size"
//	@Success		200			{object}	response.BaseResponse{data=[]response.AchievementResponse}
//	@Failure		400			{object}	response.BaseResponse
//	@Failure		500			{object}	response.BaseResponse
//	@Router			/achievements [get]
func (controller *AchievementController) GetAllAchievements(ctx *gin.Context) {
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

	achievements, err := controller.achievementService.GetAll(pageInt, pageSizeInt)
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    500,
			Status:  "Error",
			Message: "Failed to get achievements",
			Data:    nil,
		}

		ctx.JSON(500, errorResponse)
		return
	}

	webResponse := response.BaseResponse{
		Code:    200,
		Status:  "Success",
		Message: "Achievements retrieved successfully",
		Data:    achievements,
	}

	ctx.JSON(200, webResponse)
}

// GetAchievementByID godoc
//
//	@Summary		Get an achievement by ID
//	@Description	Get an achievement by ID
//	@Tags			Achievement
//	@Accept			json
//	@Produce		json
//	@Param			achievementID	path		int	true	"Achievement ID"
//	@Success		200				{object}	response.BaseResponse
//	@Failure		400				{object}	response.BaseResponse
//	@Failure		500				{object}	response.BaseResponse
//	@Router			/achievements/{achievementID} [get]
func (controller *AchievementController) GetAchievementByID(ctx *gin.Context) {
	achivementID := ctx.Param("achievementID")

	achivementIDInt, err := strconv.Atoi(achivementID)
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    400,
			Status:  "Error",
			Message: "Invalid achievementID",
			Data:    nil,
		}

		ctx.JSON(400, errorResponse)
		return
	}

	achievement, err := controller.achievementService.GetByID(uint(achivementIDInt))
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    500,
			Status:  "Error",
			Message: "Failed to get achievement",
			Data:    nil,
		}

		ctx.JSON(500, errorResponse)
		return
	}

	webResponse := response.BaseResponse{
		Code:    200,
		Status:  "Success",
		Message: "Achievement retrieved successfully",
		Data:    achievement,
	}

	ctx.JSON(200, webResponse)
}

// UpdateAchievement godoc
//
//	@Summary		Update an achievement
//	@Description	Update an achievement with the input payload
//	@Tags			Achievement
//	@Accept			json
//	@Produce		json
//	@Param			achievementID	path		int									true	"Achievement ID"
//	@Param			request			body		request.UpdateAchievementRequest	true	"Update Achievement Request"
//	@Success		200				{object}	response.BaseResponse
//	@Failure		400				{object}	response.BaseResponse
//	@Failure		500				{object}	response.BaseResponse
//	@Router			/achievements/{achievementID} [put]
func (controller *AchievementController) UpdateAchievement(ctx *gin.Context) {
	achievementID := ctx.Param("achievementID")
	updateAchievementRequest := request.UpdateAchievementRequest{}

	err := ctx.ShouldBindJSON(&updateAchievementRequest)
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

	achievementIDInt, err := strconv.Atoi(achievementID)
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    400,
			Status:  "Error",
			Message: "Invalid achievementID",
			Data:    nil,
		}

		ctx.JSON(400, errorResponse)
		return
	}

	err = controller.achievementService.Update(uint(achievementIDInt), updateAchievementRequest)
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    500,
			Status:  "Error",
			Message: "Failed to update achievement",
			Data:    nil,
		}

		ctx.JSON(500, errorResponse)
		return
	}

	webResponse := response.BaseResponse{
		Code:    200,
		Status:  "Success",
		Message: "Achievement updated successfully",
		Data:    nil,
	}

	ctx.JSON(200, webResponse)
}

// DeleteAchievement godoc
//
//	@Summary		Delete an achievement
//	@Description	Delete an achievement by ID
//	@Tags			Achievement
//	@Accept			json
//	@Produce		json
//	@Param			achievementID	path		int	true	"Achievement ID"
//	@Success		200				{object}	response.BaseResponse
//	@Failure		400				{object}	response.BaseResponse
//	@Failure		500				{object}	response.BaseResponse
//	@Router			/achievements/{achievementID} [delete]
func (controller *AchievementController) DeleteAchievement(ctx *gin.Context) {
	achievementID := ctx.Param("achievementID")

	achievementIDInt, err := strconv.Atoi(achievementID)
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    400,
			Status:  "Error",
			Message: "Invalid achievementID",
			Data:    nil,
		}

		ctx.JSON(400, errorResponse)
		return
	}

	err = controller.achievementService.Delete(uint(achievementIDInt))
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    500,
			Status:  "Error",
			Message: "Failed to delete achievement",
			Data:    nil,
		}

		ctx.JSON(500, errorResponse)
		return
	}

	webResponse := response.BaseResponse{
		Code:    200,
		Status:  "Success",
		Message: "Achievement deleted successfully",
		Data:    nil,
	}

	ctx.JSON(200, webResponse)
}

// GetAchievementWithPlayers godoc
//
//	@Summary		Get an achievement with players
//	@Description	Get an achievement with players by ID
//	@Tags			Achievement
//	@Accept			json
//	@Produce		json
//	@Param			achievementID	path		int	true	"Achievement ID"
//	@Success		200				{object}	response.BaseResponse
//	@Failure		400				{object}	response.BaseResponse
//	@Failure		500				{object}	response.BaseResponse
//	@Router			/achievements/{achievementID}/players [get]
func (controller *AchievementController) GetAchievementWithPlayers(ctx *gin.Context) {
	achievementID := ctx.Param("achievementID")

	achievementIDInt, err := strconv.Atoi(achievementID)
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    400,
			Status:  "Error",
			Message: "Invalid achievementID",
			Data:    nil,
		}

		ctx.JSON(400, errorResponse)
		return
	}

	achievement, err := controller.achievementService.GetAchievementWithPlayers(uint(achievementIDInt))
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    500,
			Status:  "Error",
			Message: "Failed to get achievement with players",
			Data:    nil,
		}

		ctx.JSON(500, errorResponse)
		return
	}

	webResponse := response.BaseResponse{
		Code:    200,
		Status:  "Success",
		Message: "Achievement with players retrieved successfully",
		Data:    achievement,
	}

	ctx.JSON(200, webResponse)
}
