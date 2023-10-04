package users

import (
	"mime/multipart"
	"time"
)

type UserCore struct {
	ID           uint
	Name         string
	Email        string
	PhoneNumber  string
	Password     string
	Address      string
	ProfilePhoto string
	UploadKTP    string
	Role         string
	MembershipID uint
	CreatedAt    time.Time
	DeletedAt    time.Time
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDataInterface interface {
	Login(email, password string) (UserCore, error)
	Insert(input UserCore, fileImages multipart.File, fileID multipart.File, filenameImages string, filenameID string) error
	SelectById(id uint) (UserCore, error)
	Delete(id uint) error
	UpdateUser(id uint, input UserCore, fileImages multipart.File, fileID multipart.File, filenameImages string, filenameID string) error
}

type UserServiceInterface interface {
	Login(email, password string) (UserCore, string, error)
	Add(input UserCore, fileImages multipart.File, fileID multipart.File, filenameImages string, filenameID string) error
	GetUserById(id uint) (UserCore, error)
	Deletebyid(id uint) error
	Update(id uint, input UserCore, fileImages multipart.File, fileID multipart.File, filenameImages string, filenameID string) error
}
