package data

import "Hannon-app/features/users"

type User struct {
	ID             uint   `gorm:"column:id;primaryKey"`
	Name           string `gorm:"name;not null"`
	UserName       string `gorm:"user_name;not null"`
	Email          string `gorm:"email;not null"`
	PhoneNumber    string `gorm:"phone_number;not null"`
	Password       string `gorm:"password;not null"`
	Address        string `gorm:"address;not null"`
	ProfilePhoto   string `gorm:"column:profile_photo"`
	UploadKTPPhoto string `gorm:"column:ktp_photo"`
	Membership     bool   `gorm:"column:Membership"`
}

func UserCoreToModel(input users.UserCore) User {
	var userModel = User{
		ID:             input.ID,
		Name:           input.Name,
		UserName:       input.UserName,
		Email:          input.Email,
		PhoneNumber:    input.PhoneNumber,
		Password:       input.Password,
		Address:        input.Address,
		ProfilePhoto:   input.ProfilePhoto,
		UploadKTPPhoto: input.UploadKTPPhoto,
		Membership:     false,
	}
	return userModel
}

func ModelToCore(input User) users.UserCore {
	var userCore = users.UserCore{
		ID:             input.ID,
		Name:           input.Name,
		UserName:       input.UserName,
		Email:          input.Email,
		PhoneNumber:    input.PhoneNumber,
		Password:       input.Password,
		Address:        input.Address,
		ProfilePhoto:   input.ProfilePhoto,
		UploadKTPPhoto: input.UploadKTPPhoto,
		Membership:     false,
	}
	return userCore
}
