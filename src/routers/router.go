package routers

import (
	"github.com/dieg0code/player-profile/src/controllers"
	"github.com/dieg0code/player-profile/src/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(authController *controllers.AuthController, userController *controllers.UserController, playerController *controllers.PlayerProfileController, achievementController *controllers.AchievementController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, "API is running")
	})

	baseRouter := router.Group("/api/v1")
	userRouter := baseRouter.Group("/users")
	playerRouter := baseRouter.Group("/players")
	achievementRouter := baseRouter.Group("/achievements")

	// Public routes
	userRouter.POST("", userController.CreateUser)
	baseRouter.POST("/login", authController.Login)

	// Apply JWTAuthMiddleware to routes that require authentication
	userRouter.Use(middleware.JWTAuthMiddleware())
	playerRouter.Use(middleware.JWTAuthMiddleware())
	achievementRouter.Use(middleware.JWTAuthMiddleware())

	// User routes
	userRouter.GET("", userController.GetAllUsers)
	userRouter.GET("/:userID", userController.GetUserByID)
	userRouter.PUT("/:userID", middleware.RoleCheckUsersMiddleware(), userController.UpdateUser)
	userRouter.DELETE("/:userID", middleware.RoleCheckUsersMiddleware(), userController.DeleteUser)

	// Player routes
	playerRouter.POST("", playerController.CreatePlayerProfile)
	playerRouter.GET("", playerController.GetAllPlayers)
	playerRouter.GET("/:playerID", playerController.GetPlayerByID)
	playerRouter.PUT("/:playerID", middleware.RoleCheckPlayersMiddleware(playerController.GetPlayerByIDFromService), playerController.UpdatePlayer)
	playerRouter.DELETE("/:playerID", middleware.RoleCheckPlayersMiddleware(playerController.GetPlayerByIDFromService), playerController.DeletePlayer)

	// Achievement routes
	achievementRouter.POST("", middleware.AuthorizationAchievementMiddleware(), achievementController.CreateAchievement)
	achievementRouter.GET("", achievementController.GetAllAchievements)
	achievementRouter.GET("/:achievementID", achievementController.GetAchievementByID)
	achievementRouter.PUT("/:achievementID", middleware.AuthorizationAchievementMiddleware(), achievementController.UpdateAchievement)
	achievementRouter.DELETE("/:achievementID", middleware.AuthorizationAchievementMiddleware(), achievementController.DeleteAchievement)

	return router
}
