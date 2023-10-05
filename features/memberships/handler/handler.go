package handler

import (
	"Hannon-app/features/memberships"
)

type MembershipHandler struct {
	membershipService memberships.MembershipServiceInterface
}

func New(service memberships.MembershipServiceInterface) *MembershipHandler {
	return &MembershipHandler{
		membershipService: service,
	}
}
