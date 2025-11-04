package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/models"
)

var users = []models.User{
	{Id: 1, Name: "Yoga"},
	{Id: 2, Name: "Dimas"},
}

type Authentication struct{}

type UserAuthentication interface{
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

func AuthenticationAcoount(ua UserAuthentication) UserAuthentication{
	return ua
}

func(a *Authentication) Register(ctx *gin.Context) {
		var form models.User
		if err := ctx.ShouldBind(&form); err != nil {
			ctx.JSON(400, models.Response{
				Success: false,
				Message: "Gagal membaca form data",
				Data:    nil,
			})
			return
		}

		if form.Name == "" || form.Email == "" || form.Password == "" {
			ctx.JSON(400, models.Response{
				Success: false,
				Message: "Nama, Email, Password wajib diisi",
				Data:    nil,
			})
			return
		}

		for _, u := range users {
			if u.Email == form.Email {
				ctx.JSON(400, models.Response{
					Success: false,
					Message: "Email sudah digunakan",
					Data:    nil,
				})
				return
			}
		}

		form.Id = len(users) + 1
		users = append(users, form)

		ctx.JSON(200, models.Response{
			Success: true,
			Message: "Register Berhasil",
			Data:    form,
		})
	}

	func(a *Authentication) Login(ctx *gin.Context) {
		var form models.User
		if err := ctx.ShouldBind(&form); err != nil {
			ctx.JSON(400, models.Response{
				Success: false,
				Message: "Gagal membaca form data",
				Data:    nil,
			})
			return
		}
		if form.Email == "" || form.Password == "" {
			ctx.JSON(400, models.Response{
				Success: false,
				Message: "Password atau Email tidak boleh kosong",
			})
			return
		}

		for _, u := range users {
			if u.Email == form.Email && u.Password == form.Password {
				ctx.JSON(200, models.Response{
					Success: true,
					Message: "Login Berhasil",
					Data:    u,
				})
				return
			}
		}
		ctx.JSON(400, models.Response{
			Success: false,
			Message: "Email atau Password salah",
			Data:    nil,
		})
	}