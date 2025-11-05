package controller

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/matthewhartstonge/argon2"
	"main.go/models"
)

var users = []models.User{
	{Id: 1, Name: "Yoga"},
	{Id: 2, Name: "Dimas"},
}

type Authentication struct{}

type UserAuthentication interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

func AuthenticationAcoount(ua UserAuthentication) UserAuthentication {
	return ua
}

func CorsMiddleware(r *gin.Engine) gin.HandlerFunc {
	godotenv.Load()
	env := os.Getenv("ORIGIN_URL")
	
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", env)
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		ctx.Next()
	}
}



func (a *Authentication) Register(ctx *gin.Context) {
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

	argon := argon2.DefaultConfig()

	encoded, err := argon.HashEncoded([]byte(form.Password))
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Gagal hash password"})
		return
	}
	form.Password = string(encoded)
	form.Id = len(users) + 1

	users = append(users, form)

	ctx.JSON(200, models.Response{
		Success: true,
		Message: "Register Berhasil",
		Data:    form,
	})
}

func (a *Authentication) Login(ctx *gin.Context) {
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
		if u.Email == form.Email {
			ok, err := argon2.VerifyEncoded([]byte(form.Password), []byte(u.Password))
			if err != nil {
				ctx.JSON(400, models.Response{
					Success: false,
					Message: "Gagal verifikasi password",
					Data:    nil,
				})
				return
			}

			if !ok {
				ctx.JSON(400, models.Response{
					Success: false,
					Message: "Password salah",
					Data:    nil,
				})
				return
			}

			ctx.JSON(200, models.Response{
				Success: true,
				Message: "Login berhasil",
				Data: gin.H{
					"id":    u.Id,
					"name":  u.Name,
					"email": u.Email,
				},
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
