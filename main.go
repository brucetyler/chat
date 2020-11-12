package main

import (
  "github.com/gin-gonic/gin"
	"chat/controller"
)

func main() {
    r := gin.Default()
    r.POST("/toLogin",controller.Userlogin)
    controller.Connect();
    //监听端口默认为8080
    r.Run(":8000")
}                                       