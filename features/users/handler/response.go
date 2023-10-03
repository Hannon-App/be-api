package handler

import "Hannon-app/features/users"

type LoginResponse struct {
	ID    uint   `json:"id,omitempty"`
	Token string `json:"token,omitempty"`
}

type UserResponse struct {
	ID           uint   `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Email        string `json:"email,omitempty"`
	PhoneNumber  string `json:"phone_number,omitempty"`
	Password     string `json:"password,omitempty"`
	Address      string `json:"address,omitempty"`
	ProfilePhoto string `json:"profil_photo,omitempty"`
	UploadKTP    string `json:"image,omitempty"`
}

func UserCoreToResponse(input users.UserCore) UserResponse {
	return UserResponse{
		ID:           input.ID,
		Name:         input.Name,
		Email:        input.Email,
		PhoneNumber:  input.PhoneNumber,
		Password:     input.Password,
		Address:      input.Address,
		ProfilePhoto: input.ProfilePhoto,
		UploadKTP:    input.UploadKTP,
	}
}
