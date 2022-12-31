package user

import (
	"errors"
	"unicode/utf8"
)

type FirstName string

func NewFirstName(v string) (*FirstName, error) {
	len := utf8.RuneCountInString(v)
	if len <= 0 {
		return nil, errors.New("must be more than 1 characters")
	}
	if len > 8 {
		return nil, errors.New("must be less than 8 characters")
	}
	i := FirstName(v)
	return &i, nil
}

func (i *FirstName) String() string {
	return string(*i)
}
