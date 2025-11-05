package view

import (
	"github.com/gin-gonic/gin"
	"main.go/controller"
)

func UserAccount(r *gin.Engine){
	var ua controller.Authentication

	r.POST("/auth/register",  ua.Register) 
	r.POST("/auth/login", ua.Login)
}