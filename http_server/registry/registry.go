package registry

import (
	"github.com/aokuyama/go-study-layered/packages/app/controller"
	"github.com/aokuyama/go-study-layered/packages/app/usecase"
	"github.com/aokuyama/go-study-layered/packages/app/usecase/get_user_info"
	"github.com/aokuyama/go-study-layered/packages/infra/database/dummy"
)

type Registry struct {
}

func (r Registry) UserController(p1 usecase.GetUserInfoPresenter) *controller.UserController {
	u := dummy.UserRepositoryDummy{}
	i := get_user_info.NewGetUserInfoInteractor(p1, u)
	c := controller.NewUserController(i)
	return &c
}
