package api

type AdminIdInput struct {
	UnixID string `uri:"admin_id" binding:"required"`
}

type VerifyTokenAdminInput struct {
	Token string `json:"token" binding:"required"`
}

type AdminId struct {
	UnixAdmin string
}
