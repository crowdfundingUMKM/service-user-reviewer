package core

type UserReviewerFormatter struct {
	ID                    int    `json:"id"`
	UnixID                string `json:"unix_id"`
	Name                  string `json:"name"`
	Email                 string `json:"email"`
	EducationalBackground string `json:"educational_background"`
	Phone                 string `json:"phone"`
	BioUser               string `json:"bio_user"`
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
		BioUser:               user.BioUser,
		Token:                 token,
		StatusAccount:         user.StatusAccount,
	}
	return formatter
}

type UserDetailFormatter struct {
	ID                    int    `json:"id"`
	UnixID                string `json:"unix_id"`
	Name                  string `json:"name"`
	Email                 string `json:"email"`
	Phone                 string `json:"phone"`
	EducationalBackground string `json:"educational_background"`
	BioUser               string `json:"bio_user"`
	Address               string `json:"address"`
	Country               string `json:"country"`
	FBLink                string `json:"fb_link"`
	IGLink                string `json:"ig_link"`
	LinkedLink            string `json:"linked_link"`
	Token                 string `json:"token"`
	StatusAccount         string `json:"status_account"`
	AvatarFileName        string `json:"avatar_file_name"`
}

func FormatterUserDetail(user User, updatedUser User) UserDetailFormatter {
	formatter := UserDetailFormatter{
		ID:                    user.ID,
		UnixID:                user.UnixID,
		Name:                  user.Name,
		Email:                 user.Email,
		Phone:                 user.Phone,
		EducationalBackground: user.EducationalBackground,
		BioUser:               user.BioUser,
		Address:               user.Address,
		Country:               user.Country,
		FBLink:                user.FBLink,
		IGLink:                user.IGLink,
		LinkedLink:            user.LinkedLink,
		StatusAccount:         user.StatusAccount,
		AvatarFileName:        user.AvatarFileName,
	}
	// read data before update if null use old data
	if updatedUser.Name != "" {
		formatter.Name = updatedUser.Name
	}
	if updatedUser.Phone != "" {
		formatter.Phone = updatedUser.Phone
	}
	if updatedUser.EducationalBackground != "" {
		formatter.EducationalBackground = updatedUser.EducationalBackground
	}
	if updatedUser.BioUser != "" {
		formatter.BioUser = updatedUser.BioUser
	}
	if updatedUser.AvatarFileName != "" {
		formatter.AvatarFileName = updatedUser.AvatarFileName
	}
	if updatedUser.StatusAccount != "" {
		formatter.StatusAccount = updatedUser.StatusAccount
	}
	if updatedUser.Address != "" {
		formatter.Address = updatedUser.Address
	}
	if updatedUser.Country != "" {
		formatter.Country = updatedUser.Country
	}
	if updatedUser.FBLink != "" {
		formatter.FBLink = updatedUser.FBLink
	}
	if updatedUser.IGLink != "" {
		formatter.IGLink = updatedUser.IGLink
	}
	if updatedUser.LinkedLink != "" {
		formatter.LinkedLink = updatedUser.LinkedLink
	}
	return formatter
}

// get user admin status
type UserReviewer struct {
	UnixReviewer       string `json:"unix_reviewer"`
	StatusAccountAdmin string `json:"status_account_reviewer"`
}

// get user reviewer status
func FormatterUserReviewerID(user User) UserReviewer {
	formatter := UserReviewer{
		UnixReviewer:       user.UnixID,
		StatusAccountAdmin: user.StatusAccount,
	}
	return formatter
}

// Notif
// notify formater
type NotifyFormatter struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	TypeError   string `json:"type_error"`
	StatusNotif int    `json:"status_notif"`
}

func FormatterNotify(notify NotifReviewer) NotifyFormatter {
	formatter := NotifyFormatter{
		ID:          notify.ID,
		Title:       notify.Title,
		Description: notify.Description,
		TypeError:   notify.TypeError,
		StatusNotif: notify.StatusNotif,
	}
	return formatter
}
