package core

type DeactiveUserInput struct {
	UnixID string `json:"unix_id" binding:"required"`
}

type ActiveUserInput struct {
	UnixID string `json:"unix_id" binding:"required"`
}

type GetUserIdInput struct {
	UnixID string `uri:"unix_id" binding:"required"`
}

type RegisterUserInput struct {
	Name                  string `json:"name" binding:"required"`
	Email                 string `json:"email" binding:"required,email"`
	EducationalBackground string `json:"educational_background" binding:"required"`
	Phone                 string `json:"phone" binding:"required"`
	BioUser               string `json:"bio_user" binding:"required"`
	Password              string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

type CheckPhoneInput struct {
	Phone string `json:"phone" binding:"required"`
}

type UpdateUserInput struct {
	Name                  string `json:"name" `
	Phone                 string `json:"phone" `
	EducationalBackground string `json:"educational_background" `
	BioUser               string `json:"bio_user" `
	Address               string `json:"address" `
	Country               string `json:"country" `
	FBLink                string `json:"fb_link" `
	IGLink                string `json:"ig_link" `
	LinkedLink            string `json:"linked_link" `
}

type UpdatePasswordInput struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}
