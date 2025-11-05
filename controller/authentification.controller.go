package controller

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/matthewhartstonge/argon2"
	"main.go/models"
)

var users = []models.User{
	{Id: 1, Name: "Yoga", Email: "yoga@mail.com"},
	{Id: 2, Name: "Dimas", Email: "dimas@mail.com"},
	{Id: 3, Name: "Fiki", Email: "Fiki@mail.com"},
	{Id: 4, Name: "Sidiq", Email: "Sidiq@mail.com"},
	{Id: 5, Name: "Yoga", Email: "yoga@mail.com"},
	{Id: 6, Name: "Dimas", Email: "dimas@mail.com"},
}

type Authentication struct{}

type UserAuthentication interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

func AuthenticationAccount(ua UserAuthentication) UserAuthentication {
	return ua
}

func CorsMiddleware(r *gin.Engine) gin.HandlerFunc {
	_ = godotenv.Load()
	env := os.Getenv("ORIGIN_URL")

	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", env)
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		ctx.Next()
	}
}

// Register godoc
// @Summary Register akun baru
// @Description Mendaftarkan akun user baru via form input
// @Tags Authentication
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "Nama user"
// @Param email formData string true "Email user"
// @Param password formData string true "Password user"
// @Success 200 {object} models.Response "Register berhasil"
// @Failure 400 {object} models.Response "Gagal register"
// @Router /auth/register [post]
func (a *Authentication) Register(ctx *gin.Context) {
	var form struct {
		Name     string `form:"name"`
		Email    string `form:"email"`
		Password string `form:"password"`
	}

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(400, models.Response{
			Success: false,
			Message: "Gagal membaca data form",
		})
		return
	}

	if form.Name == "" || form.Email == "" || form.Password == "" {
		ctx.JSON(400, models.Response{
			Success: false,
			Message: "Nama, Email, Password wajib diisi",
		})
		return
	}

	for _, u := range users {
		if u.Email == form.Email {
			ctx.JSON(400, models.Response{
				Success: false,
				Message: "Email sudah digunakan",
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
	users = append(users, models.User{
		Id:       len(users) + 1,
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	})

	ctx.JSON(200, models.Response{
		Success: true,
		Message: "Register Berhasil",
		Data:    form,
	})
}

// Login godoc
// @Summary Login user
// @Description Melakukan autentikasi user via form input
// @Tags Authentication
// @Accept multipart/form-data
// @Produce json
// @Param email formData string true "Email user"
// @Param password formData string true "Password user"
// @Success 200 {object} models.Response "Login berhasil"
// @Failure 400 {object} models.Response "Login gagal"
// @Router /auth/login [post]
func (a *Authentication) Login(ctx *gin.Context) {
	var form models.LoginRequest

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(400, models.Response{
			Success: false,
			Message: "Gagal membaca form data",
		})
		return
	}

	if form.Email == "" || form.Password == "" {
		ctx.JSON(400, models.Response{
			Success: false,
			Message: "Email dan Password wajib diisi",
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
				})
				return
			}
			if !ok {
				ctx.JSON(400, models.Response{
					Success: false,
					Message: "Password salah",
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
	})
}
