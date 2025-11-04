package main

import (
	"github.com/gin-gonic/gin"
	"main.go/view"
)



func main() {
	r := gin.Default()
	view.InitView(r)
	r.Run(":8080")
}
