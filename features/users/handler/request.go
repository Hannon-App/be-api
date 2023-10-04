package handler

import "Hannon-app/features/users"

type UserRequest struct {
	Name         string `json:"name" form:"name"`
	Email        string `json:"email" form:"email"`
	PhoneNumber  string `json:"phone_number" form:"phone_number"`
	Password     string `json:"password" form:"password"`
	Address      string `json:"address" form:"address"`
	ProfilePhoto string `json:"profil_photo" form:"profil_photo"`
	UploadKTP    string `json:"ktp_photo" form:"ktp_photo"`
	MembershipID uint   `json:"membership_id" form:"membership_id"`
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
		Name:         user.Name,
		Email:        user.Email,
		PhoneNumber:  user.PhoneNumber,
		Password:     user.Password,
		Address:      user.Address,
		ProfilePhoto: user.ProfilePhoto,
		UploadKTP:    user.UploadKTP,
		MembershipID: user.MembershipID,
	}
}
