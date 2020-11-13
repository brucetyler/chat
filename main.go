package main

import (
  "github.com/gin-gonic/gin"
	"chat/controller"
)

func main() {
    r := gin.Default()
    r.POST("/toLogin",controller.Userlogin)
    r.GET("/ws",controller.Connect)
    r.Run(":8000")
}                                       