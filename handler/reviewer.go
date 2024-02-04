package handler

import (
	"net/http"
	"service-user-reviewer/auth"
	"service-user-reviewer/core"
	"service-user-reviewer/helper"

	"github.com/gin-gonic/gin"
)

type userReviewerHandler struct {
	userService core.Service
	authService auth.Service
}

func NewUserHandler(userService core.Service, authService auth.Service) *userReviewerHandler {
	return &userReviewerHandler{userService, authService}
}

func (h *userReviewerHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input dari user ke struct RegisterUserInput
	// struct di atas kita passing sebagai parameter service

	var input core.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// generate token
	token, err := h.authService.GenerateToken(newUser.UnixID)
	if err != nil {
		if err != nil {
			response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	formatter := core.FormatterUser(newUser, token)

	if formatter.StatusAccount == "active" {
		response := helper.APIResponse("Account has been registered and active", http.StatusOK, "success", formatter)
		c.JSON(http.StatusOK, response)
		return
	}

	data := gin.H{
		"status": "Account has been registered, but you must wait admin to active your account",
	}

	response := helper.APIResponse("Account has been registered but you must wait admin or review to active your account", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}

func (h *userReviewerHandler) Login(c *gin.Context) {

	var input core.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	// generate token
	token, err := h.authService.GenerateToken(loggedinUser.UnixID)

	// check if account deactive not save token
	if loggedinUser.StatusAccount == "deactive" {
		errorMessage := gin.H{"errors": "Your account is deactive by admin"}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// save toke to database
	_, err = h.userService.SaveToken(loggedinUser.UnixID, token)

	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// end save token to database

	if err != nil {
		if err != nil {
			response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	// check role acvtive and not send massage your account deactive
	// if loggedinUser.StatusAccount == "deactive" {
	// 	errorMessage := gin.H{"errors": "Your account is deactive by admin"}
	// 	response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
	// 	c.JSON(http.StatusUnprocessableEntity, response)
	// 	return
	// }
	formatter := core.FormatterUser(loggedinUser, token)

	response := helper.APIResponse("Succesfuly loggedin", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userReviewerHandler) CheckEmailAvailability(c *gin.Context) {
	var input core.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *userReviewerHandler) CheckPhoneAvailability(c *gin.Context) {
	var input core.CheckPhoneInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Phone checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isPhoneAvailable, err := h.userService.IsPhoneAvailable(input)

	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := helper.APIResponse("Phone checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isPhoneAvailable,
	}

	metaMessage := "Phone has been registered"

	if isPhoneAvailable {
		metaMessage = "Phone is available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *userReviewerHandler) GetUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(core.User)

	// check f account deactive
	if currentUser.StatusAccount == "deactive" {
		errorMessage := gin.H{"errors": "Your account is deactive by admin"}
		response := helper.APIResponse("Get user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	// if you logout you can't get user
	if currentUser.Token == "" {
		errorMessage := gin.H{"errors": "Your account is logout"}
		response := helper.APIResponse("Get user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := core.FormatterUserDetail(currentUser, currentUser)

	response := helper.APIResponse("Successfuly get user by middleware", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userReviewerHandler) UpdateUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(core.User)

	// if account deactive
	if currentUser.StatusAccount == "deactive" {
		errorMessage := gin.H{"errors": "Your account is deactive by admin"}
		response := helper.APIResponse("Update user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// if you logout you can't get user
	if currentUser.Token == "" {
		errorMessage := gin.H{"errors": "Your account is logout"}
		response := helper.APIResponse("Get user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var inputData core.UpdateUserInput

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Update user failed, input data failure", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedUser, err := h.userService.UpdateUserByUnixID(currentUser.UnixID, inputData)
	if err != nil {
		response := helper.APIResponse("Update user failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := core.FormatterUserDetail(currentUser, updatedUser)

	response := helper.APIResponse("User has been updated", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}
