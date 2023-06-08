package reviewer

import "time"

type User struct {
	ID                    int
	UReviewerID           string
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
