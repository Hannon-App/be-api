package handler

type LoginResponse struct {
	ID    uint   `json:"id"`
	Role  string `json:"role"`
	Token string `json:"token"`
}
