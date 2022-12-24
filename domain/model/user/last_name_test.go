package user_test

import (
	"testing"

	. "github.com/aokuyama/go-study-layered/domain/model/user"

	"github.com/stretchr/testify/assert"
)

func TestLastName(t *testing.T) {
	var i *LastName
	var err error
	i, err = NewLastName("1")
	assert.Equal(t, "1", i.String(), "有効な苗字")
	assert.NoError(t, err)
	i, err = NewLastName("ああああああああ")
	assert.Equal(t, "ああああああああ", i.String(), "有効な苗字")
	assert.NoError(t, err)
}

func TestErrorLastName(t *testing.T) {
	var i *LastName
	var err error
	i, err = NewLastName("")
	assert.Nil(t, i)
	assert.Error(t, err, "0文字は許容しない")
	i, err = NewLastName("ああああああああa")
	assert.Nil(t, i)
	assert.Error(t, err, "8文字超えは許容しない")
}
