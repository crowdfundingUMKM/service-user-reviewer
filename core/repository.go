package core

import "gorm.io/gorm"

// KONTRAK
type Repository interface {
	Save(user User) (User, error)

	FindByUnixID(unix_id string) (User, error)
	FindByEmail(email string) (User, error)
	FindByPhone(phone string) (User, error)
	UpdateToken(user User) (User, error)
	Update(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByUnixID(unix_id string) (User, error) {
	var user User

	err := r.db.Where("unix_id = ?", unix_id).Find(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil

}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil

}

func (r *repository) FindByPhone(phone string) (User, error) {
	var user User

	err := r.db.Where("phone = ?", phone).Find(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil

}
func (r *repository) Update(user User) (User, error) {
	err := r.db.Model(&user).Updates(User{Name: user.Name, Phone: user.Phone, BioUser: user.BioUser, EducationalBackground: user.EducationalBackground, Address: user.Address, Country: user.Country, FBLink: user.FBLink, IGLink: user.IGLink}).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) UpdateToken(user User) (User, error) {
	err := r.db.Model(&user).Update("token", user.Token).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
