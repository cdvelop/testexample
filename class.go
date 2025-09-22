package testexample

type User struct {
	name  string
	email string
}

func NewUser(name, email string) *User {
	return &User{
		name:  name,
		email: email,
	}
}

func (u *User) GetInfo() string {
	return u.name + " <" + u.email + ">"
}
