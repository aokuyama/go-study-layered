package get_user_info

import (
	"fmt"

	"github.com/aokuyama/go-study-layered/app/usecase"
)

type GetUserInfoPresenterCli struct {
}

func NewGetUserInfoPresenterCli() GetUserInfoPresenterCli {
	return GetUserInfoPresenterCli{}
}

func (p GetUserInfoPresenterCli) Render(o usecase.GetUserInfoOutput) error {
	if o.User != nil {
		fmt.Println(o.User.String())
	} else {
		fmt.Println("user not exist")
	}
	return nil
}
