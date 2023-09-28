package handler

type UserRequest struct {
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

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
