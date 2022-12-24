package user_test

import (
	"testing"

	. "github.com/aokuyama/go-study-layered/domain/model/user"

	"github.com/stretchr/testify/assert"
)

func TestEmail(t *testing.T) {
	var i *Email
	var err error
	i, err = NewEmail("ok@exmaple.com")
	assert.Equal(t, "ok@exmaple.com", i.String(), "有効なメールアドレス")
	assert.NoError(t, err)
	i, err = NewEmail("1@exmaple.com")
	assert.Equal(t, "1@exmaple.com", i.String(), "有効なメールアドレス")
	assert.NoError(t, err)
}

func TestErrorEmail(t *testing.T) {
	var i *Email
	var err error
	i, err = NewEmail("")
	assert.Nil(t, i)
	assert.Error(t, err, "不正なメールアドレス")
	i, err = NewEmail("abcdeFGHI123")
	assert.Nil(t, i)
	assert.Error(t, err, "不正なメールアドレス")
	i, err = NewEmail("ok@@exmaple.com")
	assert.Nil(t, i)
	assert.Error(t, err, "不正なメールアドレス")
	i, err = NewEmail("@exmaple.com")
	assert.Nil(t, i)
	assert.Error(t, err, "不正なメールアドレス")
}
