package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/health_check", func(c *gin.Context) {
		c.String(200, "OK")
	})
	r.Run()
}
