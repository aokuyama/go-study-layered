package user

import (
	"errors"
	"regexp"
)

type UserID string

var re_user_id *regexp.Regexp

func NewUserID(v string) (*UserID, error) {
	len := len(v)
	if len <= 0 {
		return nil, errors.New("must be more than 1 characters")
	}
	if len > 16 {
		return nil, errors.New("must be less than 16 characters")
	}
	if re_user_id == nil {
		re_user_id = regexp.MustCompile(`[^0-9a-zA-Z]`)
	}
	if re_user_id.MatchString(v) {
		return nil, errors.New("must be alphanumeric")
	}
	i := UserID(v)
	return &i, nil
}

func (i *UserID) String() string {
	return string(*i)
}
