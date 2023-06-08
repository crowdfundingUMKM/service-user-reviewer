package reviewer

type RegisterUserInput struct {
	Name                  string `json:"name" binding:"required"`
	Email                 string `json:"email" binding:"required,email"`
	EducationalBackground string `json:"educational_background" binding:"required"`
	Phone                 string `json:"phone" binding:"required"`
	Description           string `json:"description"`
	Password              string `json:"password" binding:"required"`
}
