package handler

import "Hannon-app/features/users"

type LoginResponse struct {
	ID         uint   `json:"id"`
	Membership bool   `json:"membership"`
	Token      string `json:"token"`
}

type UserResponseAll struct {
	ID          uint   `gorm:"column:id;primaryKey"`
	Name        string `gorm:"name;not null"`
	UserName    string `gorm:"user_name;not null"`
	Email       string `gorm:"email;not null"`
	PhoneNumber string `gorm:"phone_number;not null"`
	Address     string `gorm:"address;not null"`
	Membership  bool   `gorm:"column:Membership"`
}

func UserCoreToResponseAll(input users.UserCore) UserResponseAll {
	var userResp = UserResponseAll{
		ID:          input.ID,
		Name:        input.Name,
		UserName:    input.UserName,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Address:     input.Address,
		Membership:  false,
	}
	return userResp
}
