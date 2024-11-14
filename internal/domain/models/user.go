package models

import "github.com/google/uuid"

type User struct {
	id        string
	email     string
	firstName string
	lastname  string
}

func NewUser(email, firstName, lastName string) *User {
	return &User{
		email:     email,
		firstName: firstName,
		lastname:  lastName,
	}
}

func (u *User) Email() string {
	return u.email
}

func (u *User) FullName() string {
	return u.firstName + " " + u.lastname
}

func (u *User) FirstName() string {
	return u.firstName
}

func (u *User) LastName() string {
	return u.lastname
}

func (u *User) SetID() {
	id := uuid.NewString()
	u.id = id
}

func (u *User) ID() string {
	return u.id
}
