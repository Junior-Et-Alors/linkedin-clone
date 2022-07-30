package entity

import (
	"time"
)

type User struct {
	_id       int
	Email     string
	Password  []byte
	Mdp       string
	Firstname string
	Lastname  string
	Age       int
	CreatedAt time.Time
	UpdateAt  time.Time
}

func (user User) GetId() int {
	return user._id
}

func (user User) GetEmail() string {
	return user.Email
}

func (user *User) SetEmail(email string) {
	user.Email = email
}

func (user User) GetPassword() []byte {
	return user.Password
}

func (user *User) SetPassword(Password []byte) {
	user.Password = Password
}

func (user User) GetMdp() string {
	return user.Mdp
}

// func (user *User) SetPasswordHashed(PasswordHashed []byte) {
// 	user.PasswordHashed = PasswordHashed
// }

func (user User) GetFirstName() string {
	return user.Firstname
}

func (user User) GetLastname() string {
	return user.Lastname
}

func (user User) GetAge() int {
	return user.Age
}

func (user User) GetCreatedAt() time.Time {
	return user.CreatedAt
}

func (user User) UpDateAt() time.Time {
	return user.CreatedAt
}
