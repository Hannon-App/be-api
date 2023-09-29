package admins

import (
	"time"
)

type AdminCore struct {
	ID        uint
	Name      string
	Username  string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type AdminDataInterface interface {
	Login(email string, password string) (dataLogin AdminCore, err error)
}

type AdminServiceInterface interface {
	Login(email string, password string) (dataLogin AdminCore, token string, err error)
}
