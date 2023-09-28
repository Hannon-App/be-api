package users

type UserCore struct {
	ID             uint
	Name           string
	UserName       string
	Email          string
	PhoneNumber    string
	Password       string
	Address        string
	ProfilePhoto   string
	UploadKTPPhoto string
	Membership     bool
}

type UserDataInterface interface {
	Login(email, password string) (UserCore, error)
}

type UserServiceInterface interface {
	Login(email, password string) (UserCore, string, error)
}
