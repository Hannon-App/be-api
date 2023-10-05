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

type UserPengguna struct {
	gorm.Model
	Name         string
	Email        string
	PhoneNumber  string
	Password     string
	Address      string
	ProfilePhoto string
	UploadKTP    string
	Role         string
	MembershipID uint
	Memberships  Memberships
}

type Memberships struct {
	JenisMembership string `json:"jenis_membership"`
	Status          string `json:"status"`
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

func ModelToUserPengguna(input UserPengguna) users.UserCore {
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
		Membership:   MembersipToUserCore(input.Memberships),
	}
}

func MembersipToUserCore(input Memberships) users.MembershipCore {
	return users.MembershipCore{
		JenisMembership: input.JenisMembership,
		Status:          input.Status,
	}
}
