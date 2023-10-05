package handler

type TenantLoginResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Token string `json:"token"`
}


type TenantResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Images    string `json:"images"`
	OpenTime  string `json:"open_time"`
	CloseTime string `json:"close_time"`
}
