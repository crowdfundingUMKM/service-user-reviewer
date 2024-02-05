package core

import "time"

type User struct {
	ID                    int       `json:"id"`
	UnixID                string    `json:"unix_id"`
	Name                  string    `json:"name"`
	Email                 string    `json:"email"`
	Phone                 string    `json:"phone"`
	Country               string    `json:"country"`
	Address               string    `json:"address"`
	EducationalBackground string    `json:"educational_background"`
	BioUser               string    `json:"bio_user"`
	FBLink                string    `json:"fb_link"`
	IGLink                string    `json:"ig_link"`
	LinkedLink            string    `json:"linked_link"`
	PasswordHash          string    `json:"password_hash"`
	AvatarFileName        string    `json:"avatar_file_name"`
	StatusAccount         string    `json:"status_account"`
	Token                 string    `json:"token"`
	UpdateIdAdmin         string    `json:"update_id_admin"`
	UpdateAtAdmin         time.Time `json:"update_at_admin"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

type NotifReviewer struct {
	ID             int       `json:"id"`
	UserReviewerId string    `json:"user_reviewer_id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	TypeError      string    `json:"type_error"`
	Document       string    `json:"document"`
	StatusNotif    int       `json:"status_notif"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
