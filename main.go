package main

import (
	// "fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type User struct {
	Id   int    `form:"id"`
	Name string `form:"name"`
}

var users = []User{
	{Id: 1, Name: "Yoga"},
	{Id: 2, Name: "Dimas"},
}

func main() {
	r := gin.Default()
	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, Response{
			Success: true,
			Data:    users,
		})
	})

	r.GET("/users/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, _ := strconv.Atoi(idParam)
		for _, user := range users {
			if user.Id == id {
				ctx.JSON(200, Response{
					Success: true,
					Message: "User ditemukan",
					Data:    user,
				})
				return
			}
		}
		ctx.JSON(400, Response{
			Success: false,
			Message: "User tidak ditemukan",
			Data:    nil,
		})
	})

	

	r.Run(":8080")
}
