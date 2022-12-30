package main

import (
	"flag"

	"github.com/aokuyama/go-study-layered/app/controller"
	"github.com/aokuyama/go-study-layered/app/usecase/get_user_info"
	"github.com/aokuyama/go-study-layered/infra/database/dummy"
)

func main() {
	f := flag.String("user_id", "", "user_id")
	flag.Parse()

	user_id := *f

	p := get_user_info.NewGetUserInfoPresenterCli()
	r := dummy.UserRepositoryDummy{}
	get_user_info := get_user_info.NewGetUserInfoInteractor(p, r)
	c := controller.NewUserController(get_user_info)

	c.GetUserInfo(user_id)
}
