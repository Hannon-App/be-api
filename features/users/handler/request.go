package handler

import "Hannon-app/features/users"

type UserRequest struct {
	ID           uint   `gorm:"column:id;primaryKey"`
	Name         string `gorm:"name;not null"`
	UserName     string `gorm:"user_name;not null"`
	Email        string `gorm:"email;not null"`
	PhoneNumber  string `gorm:"phone_number;not null"`
	Password     string `gorm:"password;not null"`
	Address      string `gorm:"address;not null"`
	ProfilePhoto string `gorm:"column:profile_photo"`
	UploadKTP    string `gorm:"column:ktp_photo"`
	Membership   bool   `gorm:"column:Membership"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(login LoginRequest) users.Login {
	return users.Login{
		Email:    login.Email,
		Password: login.Password,
	}
}
func RequestToCore(user UserRequest) users.UserCore {
	return users.UserCore{
		ID:           user.ID,
		Name:         user.Name,
		UserName:     user.UserName,
		Email:        user.Email,
		PhoneNumber:  user.PhoneNumber,
		Password:     user.Password,
		Address:      user.Address,
		ProfilePhoto: user.ProfilePhoto,
		UploadKTP:    user.UploadKTP,
		Membership:   user.Membership,
	}
}
