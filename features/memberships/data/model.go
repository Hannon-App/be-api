package data

import (
	"Hannon-app/features/memberships"

	"gorm.io/gorm"
)

type Membership struct {
	gorm.Model
	Jenis_Membership string
	Status           string
	Start_date       string
	End_date         string
	Deskripsi        string
}

func ModelToCore(dataModel Membership) memberships.MembershipCore {
	var membershipCore = memberships.MembershipCore{
		ID:               dataModel.ID,
		Jenis_Membership: dataModel.Jenis_Membership,
		Status:           dataModel.Status,
		Start_date:       dataModel.Start_date,
		End_date:         dataModel.End_date,
		Deskripsi:        dataModel.Deskripsi,
		CreatedAt:        dataModel.CreatedAt,
		UpdatedAt:        dataModel.UpdatedAt,
	}
	return membershipCore
}

func MembershipCoreToModel(dataCore memberships.MembershipCore) Membership {
	var membershipModel = Membership{
		Model:            gorm.Model{},
		Jenis_Membership: dataCore.Jenis_Membership,
		Status:           dataCore.Status,
		Start_date:       dataCore.Start_date,
		End_date:         dataCore.End_date,
		Deskripsi:        dataCore.Deskripsi,
	}
	return membershipModel
}
