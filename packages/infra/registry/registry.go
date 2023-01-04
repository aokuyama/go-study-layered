package registry

import (
	"github.com/aokuyama/go-study-layered/packages/app/controller"
	"github.com/aokuyama/go-study-layered/packages/app/usecase"
)

type Registry interface {
	UserController(usecase.GetUserInfoPresenter) *controller.UserController
}
