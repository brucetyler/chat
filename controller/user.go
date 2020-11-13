package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"chat/model"
	"fmt"
)

type Info struct{
	Status string `json:"status"`
	Msg string `json:"msg"`
}

var user model.User

func Userlogin(c *gin.Context){
	username := c.PostForm("username");
	password := c.PostForm("password");
	fmt.Println(username,password)
	var res Info
	if username == "" || password == "" {
		res.Status = "fail";
		res.Msg = "用户名或密码不能为空"
	} else{
		user.Name = username;
		res.Status = "success";
	}
	c.JSON(http.StatusOK, res);
}