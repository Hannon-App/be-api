package handler

import "Hannon-app/features/users"

type LoginResponse struct {
	ID         uint   `json:"id"`
	Membership bool   `json:"membership"`
	Token      string `json:"token"`
}

type UserResponse struct {
	ID          uint   `gorm:"column:id;primaryKey"`
	Name        string `gorm:"name;not null"`
	UserName    string `gorm:"user_name;not null"`
	Password    string `gorm:"password;not null"`
	Email       string `gorm:"email;not null"`
	PhoneNumber string `gorm:"phone_number;not null"`
	Address     string `gorm:"address;not null"`
	UploadKTP   string `gorm:"column:ktp_photo"`
	Membership  bool   `gorm:"column:Membership"`
}

func UserCoreToResponse(input users.UserCore) UserResponse {
	return UserResponse{
		ID:          input.ID,
		Name:        input.Name,
		UserName:    input.UserName,
		Password:    input.Password,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Address:     input.Address,
		UploadKTP:   input.UploadKTP,
		Membership:  input.Membership,
	}
}
