package main

import (
	"flag"

	"github.com/aokuyama/go-study-layered/cui/console"
	"github.com/aokuyama/go-study-layered/packages/app/usecase/get_user_info"
)

func main() {
	f := flag.String("user_id", "", "user_id")
	flag.Parse()

	user_id := *f

	r := console.Registry{}
	v := console.View{}
	p1 := get_user_info.NewGetUserInfoPresenter(v)
	c := r.UserController(p1)

	err := c.GetUserInfo(user_id)
	if err != nil {
		panic(err)
	}
}
