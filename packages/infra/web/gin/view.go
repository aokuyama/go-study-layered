package gin

import (
	"net/http"

	"github.com/aokuyama/go-study-layered/packages/app/usecase"
	"github.com/gin-gonic/gin"
)

type View struct {
	Context *gin.Context
}

func (v View) UpdateGetUserInfo(m usecase.GetUserInfoViewModel) error {
	if m.IsFound {
		v.Context.JSON(http.StatusOK, gin.H{"name": m.Name})
	} else {
		v.Context.String(http.StatusNotFound, "user not found")
	}
	return nil
}
