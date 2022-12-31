package gin

import (
	"github.com/aokuyama/go-study-layered/src/app/controller"
	"github.com/aokuyama/go-study-layered/src/app/usecase/get_user_info"
	"github.com/aokuyama/go-study-layered/src/infra/database/dummy"
	"github.com/gin-gonic/gin"
)

func Server() *gin.Engine {
	r := gin.Default()
	r.GET("/health_check", func(c *gin.Context) {
		c.String(200, "OK")
	})
	r.GET("/users/:user_id", get_user_info_c)
	return r
}

func get_user_info_c(c *gin.Context) {
	user_id := c.Param("user_id")
	v := GinView{c}
	p := get_user_info.NewGetUserInfoPresenterHttp(v)
	r := dummy.UserRepositoryDummy{}
	i := get_user_info.NewGetUserInfoInteractor(p, r)
	con := controller.NewUserController(i)
	con.GetUserInfo(user_id)
}
