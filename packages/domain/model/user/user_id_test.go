package user_test

import (
	"testing"

	. "github.com/aokuyama/go-study-layered/packages/domain/model/user"

	"github.com/stretchr/testify/assert"
)

func TestUserID(t *testing.T) {
	var i *UserID
	var err error
	i, err = NewUserID("1")
	assert.Equal(t, "1", i.String(), "有効なID")
	assert.NoError(t, err)
	i, err = NewUserID("1234567890123456")
	assert.Equal(t, "1234567890123456", i.String(), "有効なID")
	assert.NoError(t, err)
	i, err = NewUserID("abcdeFGHI123")
	assert.Equal(t, "abcdeFGHI123", i.String(), "有効なID")
	assert.NoError(t, err)
}

func TestErrorUserID(t *testing.T) {
	var i *UserID
	var err error
	i, err = NewUserID("")
	assert.Nil(t, i)
	assert.Error(t, err, "0文字は許容しない")
	i, err = NewUserID("12345678901234567")
	assert.Nil(t, i)
	assert.Error(t, err, "16文字超えは許容しない")
	i, err = NewUserID("ああ12")
	assert.Nil(t, i)
	assert.Error(t, err, "半角英数以外は許容しない")
	i, err = NewUserID("_")
	assert.Nil(t, i)
	assert.Error(t, err, "半角英数以外は許容しない")
}
