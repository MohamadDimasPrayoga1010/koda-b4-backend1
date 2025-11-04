package view

import (
	"github.com/gin-gonic/gin"
	"main.go/controller"
)

func InitUser(r *gin.Engine){
	var ucer controller.User
	
	user := r.Group("users")
	user.GET("/", ucer.GetUsers)
	user.GET("/:id", ucer.GetUserId)
	user.POST("", ucer.AddUser)
	user.PATCH("/:id", ucer.EditUser)
	user.DELETE("/:id", ucer.DeleteUser)
}