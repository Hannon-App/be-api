package memberships

import (
	"time"
)

type MembershipCore struct {
	ID               uint
	Jenis_Membership string
	Status           string
	Start_date       string
	End_date         string
	Deskripsi        string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
}

type MembershipDataInterface interface {
}

type MembershipServiceInterface interface {
}
