package data

import (
	"Hannon-app/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `gorm:"name;not null"`
	Email        string `gorm:"email;not null"`
	PhoneNumber  string `gorm:"phone_number;not null"`
	Password     string `gorm:"password;not null"`
	Address      string `gorm:"address;not null"`
	ProfilePhoto string `gorm:"column:profile_photo"`
	UploadKTP    string `gorm:"column:ktp_photo"`
	Role         string `gorm:"default:user"`
	MembershipID uint   `gorm:"membership_id"`
}

func UserCoreToModel(input users.UserCore) User {
	return User{
		Model:        gorm.Model{},
		Name:         input.Name,
		Email:        input.Email,
		PhoneNumber:  input.PhoneNumber,
		Password:     input.Password,
		Address:      input.Address,
		ProfilePhoto: input.ProfilePhoto,
		UploadKTP:    input.UploadKTP,
		Role:         input.Role,
		MembershipID: input.MembershipID,
	}
}

func ModelToUserCore(input User) users.UserCore {
	return users.UserCore{
		ID:           input.ID,
		Name:         input.Name,
		Email:        input.Email,
		PhoneNumber:  input.PhoneNumber,
		Password:     input.Password,
		Address:      input.Address,
		ProfilePhoto: input.ProfilePhoto,
		UploadKTP:    input.UploadKTP,
		Role:         input.Role,
		MembershipID: input.MembershipID,
	}
}
