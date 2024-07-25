package controllers

import (
	"strconv"

	"github.com/dieg0code/player-profile/src/data/request"
	"github.com/dieg0code/player-profile/src/data/response"
	"github.com/dieg0code/player-profile/src/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

func (controller *UserController) CreateUser(ctx *gin.Context) {
	createUserRequest := request.CreateUserRequest{}

	err := ctx.ShouldBindJSON(&createUserRequest)
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

	err = controller.userService.Create(createUserRequest)
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    500,
			Status:  "Error",
			Message: "Failed to create user",
			Data:    nil,
		}

		ctx.JSON(500, errorResponse)
		return
	}

	webResponse := response.BaseResponse{
		Code:    200,
		Status:  "Success",
		Message: "User created successfully",
		Data:    nil,
	}

	ctx.JSON(200, webResponse)
}

func (controller *UserController) GetAllUsers(ctx *gin.Context) {
	page := ctx.Query("page")
	pageSize := ctx.Query("pageSize")

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

	users, err := controller.userService.GetAll(pageInt, pageSizeInt)
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    500,
			Status:  "Error",
			Message: "Failed to get users",
			Data:    nil,
		}

		ctx.JSON(500, errorResponse)
		return
	}

	webResponse := response.BaseResponse{
		Code:    200,
		Status:  "Success",
		Message: "Users fetched successfully",
		Data:    users,
	}

	ctx.JSON(200, webResponse)
}

func (controller *UserController) GetUserByID(ctx *gin.Context) {
	userID := ctx.Param("userID")

	userIDUint64, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    400,
			Status:  "Error",
			Message: "Invalid userID",
			Data:    nil,
		}

		ctx.JSON(400, errorResponse)
		return
	}

	userIDUint := uint(userIDUint64)

	user, err := controller.userService.GetByID(userIDUint)
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    500,
			Status:  "Error",
			Message: "Failed to get user",
			Data:    nil,
		}

		ctx.JSON(500, errorResponse)
		return
	}

	successResponse := response.BaseResponse{
		Code:    200,
		Status:  "Success",
		Message: "User fetched successfully",
		Data:    user,
	}

	ctx.JSON(200, successResponse)
}

func (controller *UserController) UpdateUser(ctx *gin.Context) {
	userID := ctx.Param("userID")

	userIDUint64, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    400,
			Status:  "Error",
			Message: "Invalid userID",
			Data:    nil,
		}

		ctx.JSON(400, errorResponse)
		return
	}

	userIDUint := uint(userIDUint64)

	updateUserRequest := request.UpdateUserRequest{}

	err = ctx.ShouldBindJSON(&updateUserRequest)
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

	err = controller.userService.Update(userIDUint, updateUserRequest)
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    500,
			Status:  "Error",
			Message: "Failed to update user",
			Data:    nil,
		}

		ctx.JSON(500, errorResponse)
		return
	}

	webResponse := response.BaseResponse{
		Code:    200,
		Status:  "Success",
		Message: "User updated successfully",
		Data:    nil,
	}

	ctx.JSON(200, webResponse)
}

func (controller *UserController) DeleteUser(ctx *gin.Context) {
	userID := ctx.Param("userID")

	userIDUint64, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    400,
			Status:  "Error",
			Message: "Invalid userID",
			Data:    nil,
		}

		ctx.JSON(400, errorResponse)
		return
	}

	userIDUint := uint(userIDUint64)

	err = controller.userService.Delete(userIDUint)
	if err != nil {
		errorResponse := response.BaseResponse{
			Code:    500,
			Status:  "Error",
			Message: "Failed to delete user",
			Data:    nil,
		}

		ctx.JSON(500, errorResponse)
		return
	}

	webResponse := response.BaseResponse{
		Code:    200,
		Status:  "Success",
		Message: "User deleted successfully",
		Data:    nil,
	}

	ctx.JSON(200, webResponse)
}
