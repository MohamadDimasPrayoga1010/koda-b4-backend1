package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"main.go/docs"
	"main.go/view"
)

// @title API User & Auth Example
// @version 1.0
// @description Ini adalah dokumentasi API sederhana menggunakan Gin & Swagger
// @termsOfService http://swagger.io/terms/
// @host localhost:8085
// @BasePath /
func main() {
	r := gin.Default()
	view.InitView(r)
	r.MaxMultipartMemory = 1 << 20 
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8085")
}
