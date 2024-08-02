package main

import (
	"log"
	"net/http"

	_ "github.com/dieg0code/player-profile/docs"
	"github.com/dieg0code/player-profile/src/config"
	"github.com/dieg0code/player-profile/src/controllers"
	"github.com/dieg0code/player-profile/src/models"
	repo "github.com/dieg0code/player-profile/src/repository/impl"
	"github.com/dieg0code/player-profile/src/routers"
	services "github.com/dieg0code/player-profile/src/services/impl"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Player Profile API
//	@version		1.0
//	@description	This is a simple API for managing player profiles and achievements

//	@host		localhost:8080
//	@BasePath	/api/v1

// @tag.name Auth
// @tag.name User
// @tag.name Player
// @tag.name Achievement
func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db := config.DatabaseConnection()
	validate := validator.New()

	passWordHasher := services.NewPassWordHasher()

	err = db.AutoMigrate(&models.User{}, &models.PlayerProfile{}, &models.Achievement{})
	if err != nil {
		panic(err)
	}
	// User repo
	userRepo := repo.NewUserRepositoryImpl(db)
	//Player profile repo
	playerProfileRepo := repo.NewPlayerProfileRepositoryImpl(db)
	//Achievement repo
	achievementRepo := repo.NewAchievementRepositoryImpl(db)

	// SERVICES

	// Auth service
	authService := services.NewAuthService(userRepo, passWordHasher, validate)

	// User service
	userService := services.NewUserServiceImpl(userRepo, validate, passWordHasher)

	// Player profile service
	playerProfileService := services.NewPlayerProfileServiceImpl(playerProfileRepo, validate)

	// Achievement service
	achievementService := services.NewAchievementServiceImpl(achievementRepo, validate)

	// CONTROLLERS

	// Auth controller
	authController := controllers.NewAuthController(authService)

	// User controller
	userController := controllers.NewUserController(userService)

	// Player profile controller
	playerController := controllers.NewPlayerProfileController(playerProfileService)

	// Achievement controller
	achievementController := controllers.NewAchievementController(achievementService)

	// ROUTER

	routes := routers.NewRouter(authController, userController, playerController, achievementController)

	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
