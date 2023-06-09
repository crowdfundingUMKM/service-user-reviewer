package reviewer

type UserReviewerFormatter struct {
	ID                    int    `json:"id"`
	UnixID                string `json:"unix_id"`
	Name                  string `json:"name"`
	Email                 string `json:"email"`
	EducationalBackground string `json:"educational_background"`
	Phone                 string `json:"phone"`
	Description           string `json:"description"`
	Token                 string `json:"token"`
	StatusAccount         string `json:"status_account"`
}

func FormatterUser(user User, token string) UserReviewerFormatter {
	formatter := UserReviewerFormatter{
		ID:                    user.ID,
		UnixID:                user.UnixID,
		Name:                  user.Name,
		Email:                 user.Email,
		EducationalBackground: user.EducationalBackground,
		Phone:                 user.Phone,
		Description:           user.Description,
		Token:                 token,
		StatusAccount:         user.StatusAccount,
	}
	return formatter
}
