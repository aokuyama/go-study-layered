package controller

import (
	"github.com/aokuyama/go-study-layered/packages/app/usecase"
	"github.com/aokuyama/go-study-layered/packages/domain/model/user"
)

type UserController struct {
	getUserInfo usecase.GetUserInfoUseCase
}

func NewUserController(
	get_user_info usecase.GetUserInfoUseCase,
) UserController {
	c := UserController{
		get_user_info,
	}
	return c
}

func (c UserController) GetUserInfo(user_id string) error {
	id, err := user.NewUserID(user_id)
	if err != nil {
		return err
	}
	i := usecase.GetUserInfoInput{UserID: *id}
	return c.getUserInfo.Handle(i)
}
