package users

type UserCore struct {
	ID           uint
	Name         string
	UserName     string
	Email        string
	PhoneNumber  string
	Password     string
	Address      string
	ProfilePhoto string
	UploadKTP    string
	Membership   bool
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDataInterface interface {
	Login(email, password string) (UserCore, error)
	Insert(input UserCore) error
	SelectById(id uint) (UserCore, error)
	Delete(id uint) error
}

type UserServiceInterface interface {
	Login(email, password string) (UserCore, string, error)
	Add(input UserCore) error
	GetUserById(id uint) (UserCore, error)
	Deletebyid(id uint) error
}
