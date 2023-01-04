package main

import (
	"github.com/aokuyama/go-study-layered/http_server/registry"

	"github.com/aokuyama/go-study-layered/packages/infra/web/gin"
)

func main() {
	r := registry.Registry{}
	gin.Server(r).Run()
}
