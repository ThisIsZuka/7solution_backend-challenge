package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/Q01", MainHandler)

	r.POST("/Q02", Main02Handler)

	r.GET("/Q03", Main03Handler)

	r.Run(":8080")
}
