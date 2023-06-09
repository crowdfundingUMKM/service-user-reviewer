package reviewer

import "time"

type User struct {
	ID                    int
	UserID                string
	Name                  string
	Email                 string
	EducationalBackground string
	Phone                 string
	Description           string
	PasswordHash          string
	AvatarFileName        string
	StatusAccount         string
	Token                 string
	CreatedAt             time.Time
	UpdatedAt             time.Time
}
