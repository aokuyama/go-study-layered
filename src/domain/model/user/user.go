package user

import (
	"encoding/json"
	"errors"
)

type User struct {
	ID        *UserID   `json:"id"`
	FirstName FirstName `json:"first_name"`
	LastName  LastName  `json:"last_name"`
	Email     Email     `json:"email"`
}

func New(
	first_name string,
	last_name string,
	email string,
) (*User, error) {
	f, err2 := NewFirstName(first_name)
	l, err3 := NewLastName(last_name)
	m, err4 := NewEmail(email)

	err := errors.Join(err2, err3, err4)
	if err != nil {
		return nil, err
	}

	u := User{
		nil,
		*f,
		*l,
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
	f, err2 := NewFirstName(first_name)
	l, err3 := NewLastName(last_name)
	m, err4 := NewEmail(email)

	err := errors.Join(err1, err2, err3, err4)
	if err != nil {
		return nil, err
	}

	u := User{
		i,
		*f,
		*l,
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
