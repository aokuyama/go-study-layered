package user

import (
	"errors"
	"unicode/utf8"
)

type Name struct {
	firstName string
	lastName  string
}

func NewName(firstName string, lastName string) (*Name, error) {
	err1 := validateName(firstName)
	err2 := validateName(lastName)
	err := errors.Join(err1, err2)
	if err != nil {
		return nil, err
	}
	n := Name{firstName, lastName}
	return &n, nil
}

func validateName(v string) error {
	len := utf8.RuneCountInString(v)
	if len <= 0 {
		return errors.New("must be more than 1 characters")
	}
	if len > 8 {
		return errors.New("must be less than 8 characters")
	}
	return nil
}

func (n *Name) String() string {
	return n.FullName()
}

func (n *Name) FullName() string {
	return n.lastName + " " + n.firstName
}
