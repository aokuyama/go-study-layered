package user

import (
	"errors"
	"unicode/utf8"
)

type LastName string

func NewLastName(v string) (*LastName, error) {
	len := utf8.RuneCountInString(v)
	if len <= 0 {
		return nil, errors.New("must be more than 1 characters")
	}
	if len > 8 {
		return nil, errors.New("must be less than 8 characters")
	}
	i := LastName(v)
	return &i, nil
}

func (i *LastName) String() string {
	return string(*i)
}
