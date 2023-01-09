package user

import (
	"encoding/json"
	"errors"
)

type User struct {
	id    *UserID
	name  Name
	email Email
}

func New(
	first_name string,
	last_name string,
	email string,
) (*User, error) {
	n, err2 := NewName(first_name, last_name)
	m, err3 := NewEmail(email)

	err := errors.Join(err2, err3)
	if err != nil {
		return nil, err
	}

	u := User{
		nil,
		*n,
		*m,
	}
	return &u, nil
}

func Load(
	id string,
	first_name string,
	last_name string,
	email string,
) (*User, error) {
	i, err1 := NewUserID(id)
	n, err2 := NewName(first_name, last_name)
	m, err3 := NewEmail(email)

	err := errors.Join(err1, err2, err3)
	if err != nil {
		return nil, err
	}

	u := User{
		i,
		*n,
		*m,
	}
	return &u, nil
}

func (u *User) String() string {
	j, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	return string(j)
}

func (u *User) FullName() string {
	return u.name.FullName()
}

func (u User) MarshalJSON() ([]byte, error) {
	var i string
	if u.id != nil {
		i = u.id.String()
	}
	v, err := json.Marshal(&struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}{
		ID:    i,
		Name:  u.name.FullName(),
		Email: u.email.String(),
	})
	return v, err
}
