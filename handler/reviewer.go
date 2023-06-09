package handler

import (
	"net/http"
	"service-user-reviewer/auth"
	"service-user-reviewer/helper"
	"service-user-reviewer/reviewer"

	"github.com/gin-gonic/gin"
)

type userReviewerHandler struct {
	userService reviewer.Service
	authService auth.Service
}

func NewUserHandler(userService reviewer.Service, authService auth.Service) *userReviewerHandler {
	return &userReviewerHandler{userService, authService}
}

func (h *userReviewerHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input dari user ke struct RegisterUserInput
	// struct di atas kita passing sebagai parameter service

	var input reviewer.RegisterUserInput

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
	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		if err != nil {
			response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	formatter := reviewer.FormatterUser(newUser, token)

	response := helper.APIResponse("Account has been registered but you must wait admin or review to active your account", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
