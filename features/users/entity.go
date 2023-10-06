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
	Membership   MembershipCore
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MembershipCore struct {
	JenisMembership string `json:"jenis_membership"`
	Status          string `json:"status"`
}

type UserDataInterface interface {
	Login(email, password string) (UserCore, error)
	Insert(input UserCore, fileImages multipart.File, fileID multipart.File, filenameImages string, filenameID string) error
	SelectById(id uint) (UserCore, error)
	Delete(id uint) error
	UpdateUser(id uint, input UserCore, fileImages multipart.File, fileID multipart.File, filenameImages string, filenameID string) error
	ReadAll(page uint, userPerPage uint, searchName string) ([]UserCore, int64, error)
}

type UserServiceInterface interface {
	Login(email, password string) (UserCore, string, error)
	Add(input UserCore, fileImages multipart.File, fileID multipart.File, filenameImages string, filenameID string) error
	GetUserById(id uint) (UserCore, error)
	Deletebyid(id uint) error
	Update(id uint, input UserCore, fileImages multipart.File, fileID multipart.File, filenameImages string, filenameID string) error
	GetAll(page, userPerPage uint, searchName string) ([]UserCore, bool, error)
}
