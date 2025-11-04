package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data any `json:"data"`
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
	r.GET("/users", func(ctx *gin.Context) {
		var u User
		if err := ctx.BindQuery(&u); err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Error binding query",
			})
			return
		}

		search := false
		for _, user := range users {
			if user.Name == u.Name {
				search = true
				break
			}
		}

		if !search {
			ctx.JSON(404, Response{
				Success: false,
				Message: fmt.Sprintf("Id %d tidak ditemukan!", u.Id),
			})
			return
		}

		ctx.JSON(200, Response{
			Success: true,
			Message: fmt.Sprintf("Halo, %s", u.Name),
		})
	})
	r.Run(":8090")
}
