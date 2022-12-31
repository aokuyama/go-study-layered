package gin

import "github.com/gin-gonic/gin"

type GinView struct {
	Context *gin.Context
}

func (v GinView) String(status int, str string) {
	v.Context.String(status, str)
}

func (v GinView) Json(status int, json any) {
	v.Context.JSON(status, json)
}
