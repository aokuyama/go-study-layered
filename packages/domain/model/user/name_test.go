package user_test

import (
	"testing"

	. "github.com/aokuyama/go-study-layered/packages/domain/model/user"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	var n *Name
	var err error
	n, err = NewName("太郎", "山田")
	assert.Equal(t, "山田 太郎", n.String(), "有効な名前")
	assert.NoError(t, err)
	n, err = NewName("Aaa", "Bbb")
	assert.Equal(t, "Bbb Aaa", n.String(), "有効な名前")
	assert.NoError(t, err)
}

func TestErrorName(t *testing.T) {
	var n *Name
	var err error
	n, err = NewName("", "太郎")
	assert.Nil(t, n)
	assert.Error(t, err, "0文字は許容しない")
	n, err = NewName("山田", "")
	assert.Nil(t, n)
	assert.Error(t, err, "0文字は許容しない")
	n, err = NewName("ああああああああa", "a")
	assert.Nil(t, n)
	assert.Error(t, err, "8文字超えは許容しない")
	n, err = NewName("a", "ああああああああa")
	assert.Nil(t, n)
	assert.Error(t, err, "8文字超えは許容しない")
}
