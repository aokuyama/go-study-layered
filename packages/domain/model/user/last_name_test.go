package user_test

import (
	"testing"

	. "github.com/aokuyama/go-study-layered/packages/domain/model/user"

	"github.com/stretchr/testify/assert"
)

func TestLastName(t *testing.T) {
	var n *LastName
	var err error
	n, err = NewLastName("1")
	assert.Equal(t, "1", n.String(), "有効な苗字")
	assert.NoError(t, err)
	n, err = NewLastName("ああああああああ")
	assert.Equal(t, "ああああああああ", n.String(), "有効な苗字")
	assert.NoError(t, err)
}

func TestErrorLastName(t *testing.T) {
	var n *LastName
	var err error
	n, err = NewLastName("")
	assert.Nil(t, n)
	assert.Error(t, err, "0文字は許容しない")
	n, err = NewLastName("ああああああああa")
	assert.Nil(t, n)
	assert.Error(t, err, "8文字超えは許容しない")
}
