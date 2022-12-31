package main

import (
	"flag"

	"github.com/aokuyama/go-study-layered/src/app/controller"
	"github.com/aokuyama/go-study-layered/src/app/usecase/get_user_info"
	"github.com/aokuyama/go-study-layered/src/infra/database/dummy"
)

func main() {
	f := flag.String("user_id", "", "user_id")
	flag.Parse()

	user_id := *f

	p := get_user_info.NewGetUserInfoPresenterCli()
	r := dummy.UserRepositoryDummy{}
	get_user_info := get_user_info.NewGetUserInfoInteractor(p, r)
	c := controller.NewUserController(get_user_info)

	err := c.GetUserInfo(user_id)
	if err != nil {
		panic(err)
	}
}
