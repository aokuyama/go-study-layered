package user_test

import (
	"testing"

	. "github.com/aokuyama/go-study-layered/packages/domain/model/user"

	"github.com/stretchr/testify/assert"
)

func TestEmail(t *testing.T) {
	var m *Email
	var err error
	m, err = NewEmail("ok@exmaple.com")
	assert.Equal(t, "ok@exmaple.com", m.String(), "有効なメールアドレス")
	assert.NoError(t, err)
	m, err = NewEmail("1@exmaple.com")
	assert.Equal(t, "1@exmaple.com", m.String(), "有効なメールアドレス")
	assert.NoError(t, err)
}

func TestErrorEmail(t *testing.T) {
	var m *Email
	var err error
	m, err = NewEmail("")
	assert.Nil(t, m)
	assert.Error(t, err, "不正なメールアドレス")
	m, err = NewEmail("abcdeFGHI123")
	assert.Nil(t, m)
	assert.Error(t, err, "不正なメールアドレス")
	m, err = NewEmail("ok@@exmaple.com")
	assert.Nil(t, m)
	assert.Error(t, err, "不正なメールアドレス")
	m, err = NewEmail("@exmaple.com")
	assert.Nil(t, m)
	assert.Error(t, err, "不正なメールアドレス")
}
