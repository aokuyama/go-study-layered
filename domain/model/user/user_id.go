package user

import (
	"errors"
	"regexp"
)

type UserID string

var re *regexp.Regexp

func NewUserID(v string) (*UserID, error) {
	if re == nil {
		re = regexp.MustCompile(`[^0-9a-zA-Z]`)
	}
	if len(v) <= 0 {
		return nil, errors.New("must be more than 1 characters")
	}
	if len(v) > 16 {
		return nil, errors.New("must be less than 16 characters")
	}
	if re.MatchString(v) {
		return nil, errors.New("must be alphanumeric")
	}
	i := UserID(v)
	return &i, nil
}

func (i *UserID) String() string {
	return string(*i)
}
