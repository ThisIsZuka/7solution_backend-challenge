package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// เรียกใช้ Engine ของ Gin
	r := gin.Default()

	// กำหนดเส้นทาง /hello ให้เรียกใช้ฟังก์ชัน helloHandler
	r.GET("/hello", MainHandler)

	// ให้ Gin รันบนพอร์ต 8080
	r.Run(":8080")
}
