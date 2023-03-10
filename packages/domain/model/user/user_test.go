package user_test

import (
	"testing"

	. "github.com/aokuyama/go-study-layered/packages/domain/model/user"
	"github.com/stretchr/testify/assert"
)

func TestEnableUser(t *testing.T) {
	var u *User
	var err error
	u, err = New("太郎", "山田", "ok@example.com")
	assert.Equal(t, `{"id":"","name":"山田 太郎","email":"ok@example.com"}`, u.String(), "有効なユーザー")
	assert.NoError(t, err)
	u, err = Load("123456789", "太郎", "山田", "ok@example.com")
	assert.Equal(t, `{"id":"123456789","name":"山田 太郎","email":"ok@example.com"}`, u.String(), "有効なユーザー")
	assert.NoError(t, err)
}

func TestErrorUser(t *testing.T) {
	var u *User
	var err error
	u, err = New("", "山田", "ok@example.com")
	assert.Nil(t, u)
	assert.Error(t, err)
	u, err = New("", "山田", "ok@example.com")
	assert.Nil(t, u)
	assert.Error(t, err)
	u, err = New("太郎", "山田", "")
	assert.Nil(t, u)
	assert.Error(t, err)
	u, err = Load("", "太郎", "山田", "ok@example.com")
	assert.Nil(t, u)
	assert.Error(t, err)
}
