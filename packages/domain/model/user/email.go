package user

import (
	"errors"
	"regexp"
)

type Email string

var re_email *regexp.Regexp

func NewEmail(v string) (*Email, error) {
	if re_email == nil {
		re_email = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	}
	if !re_email.MatchString(v) {
		return nil, errors.New("malformed")
	}
	i := Email(v)
	return &i, nil
}

func (i *Email) String() string {
	return string(*i)
}
