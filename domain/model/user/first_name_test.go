package user_test

import (
	"testing"

	. "github.com/aokuyama/go-study-layered/domain/model/user"

	"github.com/stretchr/testify/assert"
)

func TestFirstName(t *testing.T) {
	var i *FirstName
	var err error
	i, err = NewFirstName("1")
	assert.Equal(t, "1", i.String(), "有効な苗字")
	assert.NoError(t, err)
	i, err = NewFirstName("ああああああああ")
	assert.Equal(t, "ああああああああ", i.String(), "有効な苗字")
	assert.NoError(t, err)
}

func TestErrorFirstName(t *testing.T) {
	var i *FirstName
	var err error
	i, err = NewFirstName("")
	assert.Nil(t, i)
	assert.Error(t, err, "0文字は許容しない")
	i, err = NewFirstName("ああああああああa")
	assert.Nil(t, i)
	assert.Error(t, err, "8文字超えは許容しない")
}
