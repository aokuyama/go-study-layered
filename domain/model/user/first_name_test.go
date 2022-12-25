package user_test

import (
	"testing"

	. "github.com/aokuyama/go-study-layered/domain/model/user"

	"github.com/stretchr/testify/assert"
)

func TestFirstName(t *testing.T) {
	var n *FirstName
	var err error
	n, err = NewFirstName("1")
	assert.Equal(t, "1", n.String(), "有効な苗字")
	assert.NoError(t, err)
	n, err = NewFirstName("ああああああああ")
	assert.Equal(t, "ああああああああ", n.String(), "有効な苗字")
	assert.NoError(t, err)
}

func TestErrorFirstName(t *testing.T) {
	var n *FirstName
	var err error
	n, err = NewFirstName("")
	assert.Nil(t, n)
	assert.Error(t, err, "0文字は許容しない")
	n, err = NewFirstName("ああああああああa")
	assert.Nil(t, n)
	assert.Error(t, err, "8文字超えは許容しない")
}
