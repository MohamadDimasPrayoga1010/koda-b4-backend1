package view

import "github.com/gin-gonic/gin"

func InitView(r *gin.Engine) {
	InitUser(r)
	UserAccount(r)
}
