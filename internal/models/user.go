package models

type User struct {
	Login    string `json:"login" validate:"required,email"`
	Password string `json:"password"`
}

// CreateUser ..
func CreateUser(login, password string) User {
	return User{
		Login:    login,
		Password: password,
	}
}

// Validate ...
func (u *User) Validate() error {
	return validate.Struct(u)
}
