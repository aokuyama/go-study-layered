package gin

import (
	"github.com/aokuyama/go-study-layered/packages/app/usecase/get_user_info"
	"github.com/aokuyama/go-study-layered/packages/infra/registry"
	"github.com/gin-gonic/gin"
)

func Server(r registry.Registry) *gin.Engine {
	g := gin.Default()
	g.GET("/health_check", func(c *gin.Context) {
		c.String(200, "OK")
	})
	g.GET("/users/:user_id", func(c *gin.Context) {
		user_id := c.Param("user_id")
		v := View{c}
		p1 := get_user_info.NewGetUserInfoPresenter(v)
		r.UserController(p1).GetUserInfo(user_id)
	})
	return g
}
