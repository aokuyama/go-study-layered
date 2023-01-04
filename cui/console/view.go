package console

import (
	"fmt"

	"github.com/aokuyama/go-study-layered/packages/app/usecase"
)

type View struct {
}

func (v View) UpdateGetUserInfo(m usecase.GetUserInfoViewModel) error {
	if m.IsFound {
		fmt.Println("user name:", m.Name)
	} else {
		fmt.Println("user not found")
	}
	return nil
}
