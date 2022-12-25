package main

import (
	"github.com/aokuyama/go-study-layered/app/controller"
	"github.com/aokuyama/go-study-layered/app/usecase/get_user_info"
	"github.com/aokuyama/go-study-layered/domain/model/user"
)

func main() {
	p := get_user_info.NewGetUserInfoPresenterCli()
	r := UserRepositoryDummy{}
	get_user_info := get_user_info.NewGetUserInfoInteractor(p, r)
	c := controller.NewUserController(get_user_info)
	// ユーザーが存在する場合
	c.GetUserInfo("123456789")
	// ユーザーが存在しない場合
	c.GetUserInfo("12345678")
	// エラー時
	err := c.GetUserInfo("")
	println(err.Error())
}

type UserRepositoryDummy struct {
}

func (r UserRepositoryDummy) LoadByID(id user.UserID) (*user.User, error) {
	if id == "123456789" {
		return user.New("太郎", "山田", "ok@example.com")
	}
	return nil, nil
}
