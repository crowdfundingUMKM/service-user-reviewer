package handler

import (
	"net/http"
	api_admin "service-user-reviewer/api/admin"
	"service-user-reviewer/auth"
	"service-user-reviewer/core"
	"service-user-reviewer/helper"

	"github.com/gin-gonic/gin"
)

type notifHandler struct {
	userService core.Service
	authService auth.Service
}

func NewNotifHandler(userService core.Service, authService auth.Service) *notifHandler {
	return &notifHandler{userService, authService}
}

func (h *notifHandler) ReportToAdmin(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(core.User)

	// check if account is deactivated
	if currentUser.StatusAccount == "deactive" {
		errorMessage := gin.H{"errors": "Your account is deactive by admin"}
		response := helper.APIResponse("Get user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// if the user is logged out, they can't get user data
	if currentUser.Token == "" {
		errorMessage := gin.H{"errors": "Your account is logout"}
		response := helper.APIResponse("Get user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// create input
	var input core.ReportToAdminInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Invalid report to admin", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newNotif, err := h.userService.ReportAdmin(currentUser.UnixID, input)
	if err != nil {
		response := helper.APIResponse("Report to admin failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := core.FormatterNotify(newNotif)

	response := helper.APIResponse("Report to admin success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// GetNotifToAdmin
func (h *notifHandler) GetNotifToAdmin(c *gin.Context) {
	currentAdmin := c.MustGet("currentUserAdmin").(api_admin.AdminId)

	if currentAdmin.UnixAdmin == "" {
		errorMessage := gin.H{"errors": "Your account admin is logout"}
		response := helper.APIResponse("Get all reports", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	users, err := h.userService.GetAllReports()
	if err != nil {
		response := helper.APIResponse("Failed All report campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List of report campaign", http.StatusOK, "success", users)
	c.JSON(http.StatusOK, response)
	return

}
