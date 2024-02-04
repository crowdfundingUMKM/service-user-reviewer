package main

import (
	"fmt"
	"log"
	"os"
	"service-user-reviewer/auth"
	"service-user-reviewer/config"
	"service-user-reviewer/core"
	"service-user-reviewer/database"
	"service-user-reviewer/handler"
	"service-user-reviewer/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// setup log
	// config.InitLog()
	// setup repository
	db := database.NewConnectionDB()
	userReviewerRepository := core.NewRepository(db)

	// setup service
	userReviewerService := core.NewService(userReviewerRepository)
	authService := auth.NewService()

	// setup handler
	userHandler := handler.NewUserHandler(userReviewerService, authService)

	// RUN SERVICE
	router := gin.Default()

	// CORS
	// setup cors
	corsConfig := config.InitCors()
	router.Use(cors.New(corsConfig))

	// group route
	api := router.Group("api/v1")

	// Rounting admin
	// admin request -> service user admin
	api.GET("/admin/log_service_toAdmin/:admin_id", middleware.AuthApiAdminMiddleware(authService, userReviewerService), userHandler.GetLogtoAdmin)
	api.GET("/admin/service_status/:admin_id", middleware.AuthApiAdminMiddleware(authService, userReviewerService), userHandler.ServiceHealth)
	api.PUT("/admin/deactive_user/:admin_id", middleware.AuthApiAdminMiddleware(authService, userReviewerService), userHandler.DeactiveUser)
	api.PUT("/admin/active_user/:admin_id", middleware.AuthApiAdminMiddleware(authService, userReviewerService), userHandler.ActiveUser)

	// make endpoint get all user by admin
	// api.GET("/admin/get_all_user/:admin_id", middleware.AuthApiAdminMiddleware(authService, userReviewerService), userHandler.GetAllUserData)

	// route give information to user about admin
	// api.GET("/campaign/getUserReviewerID/:unix_id", userHandler.GetInfoAdminID)

	// verify token
	api.GET("/verifyTokenReviewer", middleware.AuthMiddleware(authService, userReviewerService), userHandler.VerifyToken)

	// Rounting user
	api.POST("/register_reviewer", userHandler.RegisterUser)
	api.POST("/login_reviewer", userHandler.Login)
	api.GET("/email_check", userHandler.CheckEmailAvailability)
	api.GET("/phone_check", userHandler.CheckPhoneAvailability)

	api.GET("/get_user", middleware.AuthMiddleware(authService, userReviewerService), userHandler.GetUser)

	api.PUT("/update_profile", middleware.AuthMiddleware(authService, userReviewerService), userHandler.UpdateUser)

	api.PUT("/update_password", middleware.AuthMiddleware(authService, userReviewerService), userHandler.UpdatePassword)
	//make create image profile user by unix_id this for update -> update same
	api.POST("/upload_avatar", middleware.AuthMiddleware(authService, userReviewerService), userHandler.UploadAvatar)

	// make logout user
	api.DELETE("/logout_reviewer", middleware.AuthMiddleware(authService, userReviewerService), userHandler.LogoutUser)

	// Notif route
	// api.POST("/report_to_admin", middleware.AuthMiddleware(authService, userCampaignService), notifHandler.ReportToAdmin)
	// api.GET("/admin/get_notif_admin", middleware.AuthApiAdminMiddleware(authService, userCampaignService), notifHandler.GetNotifToAdmin)

	url := fmt.Sprintf("%s:%s", os.Getenv("SERVICE_HOST"), os.Getenv("SERVICE_PORT"))
	router.Run(url)
}
