package data

import "Hannon-app/features/users"

type User struct {
	ID           uint   `gorm:"column:id;primaryKey"`
	Name         string `gorm:"name;not null"`
	UserName     string `gorm:"user_name;not null"`
	Email        string `gorm:"email;not null"`
	PhoneNumber  string `gorm:"phone_number;not null"`
	Password     string `gorm:"password;not null"`
	Address      string `gorm:"address;not null"`
	ProfilePhoto string `gorm:"column:profile_photo"`
	UploadKTP    string `gorm:"column:ktp_photo"`
	Membership   bool   `gorm:"column:membership"`
}

func UserCoreToModel(input users.UserCore) User {
	return User{
		ID:           input.ID,
		Name:         input.Name,
		UserName:     input.UserName,
		Email:        input.Email,
		PhoneNumber:  input.PhoneNumber,
		Password:     input.Password,
		Address:      input.Address,
		ProfilePhoto: input.ProfilePhoto,
		UploadKTP:    input.UploadKTP,
		Membership:   input.Membership,
	}
}

func ModelToUserCore(input User) users.UserCore {
	return users.UserCore{
		ID:           input.ID,
		Name:         input.Name,
		UserName:     input.UserName,
		Email:        input.Email,
		PhoneNumber:  input.PhoneNumber,
		Password:     input.Password,
		Address:      input.Address,
		ProfilePhoto: input.ProfilePhoto,
		UploadKTP:    input.UploadKTP,
		Membership:   input.Membership,
	}
}
