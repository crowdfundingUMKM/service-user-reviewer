package main

import (
	"fmt"
	"log"
	"os"
	"service-user-reviewer/auth"
	"service-user-reviewer/database"
	"service-user-reviewer/handler"
	"service-user-reviewer/reviewer"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// setup log
	// L.InitLog()
	// setup repository
	db := database.NewConnectionDB()
	userReviewerRepository := reviewer.NewRepository(db)

	// setup service
	userReviewerService := reviewer.NewService(userReviewerRepository)
	authService := auth.NewService()

	// setup handler
	userHandler := handler.NewUserHandler(userReviewerService, authService)

	// RUN SERVICE
	router := gin.Default()
	api := router.Group("api/v1")

	// Rounting admin
	api.POST("/register_reviewer", userHandler.RegisterUser)

	url := fmt.Sprintf("%s:%s", os.Getenv("SERVICE_HOST"), os.Getenv("SERVICE_PORT"))
	router.Run(url)
}
