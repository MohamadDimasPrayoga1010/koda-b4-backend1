package controller

import (
	"strconv"
	"github.com/matthewhartstonge/argon2"
	"github.com/gin-gonic/gin"
	"main.go/models"
)

type User struct{}

type UserController interface{
	GetUsers(ctx *gin.Context)
	GetUserId(ctx *gin.Context)
	EditUser(ctx *gin.Context)
	AddUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

func NewUserController(uc UserController) UserController{
	return uc
}

func (u *User) GetUsers(ctx *gin.Context) {
	ctx.JSON(200, models.Response{
		Success: true,
		Data:    users,
	})
}

func (u *User) GetUserId(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)
	for _, user := range users {
		if user.Id == id {
			ctx.JSON(200, models.Response{
				Success: true,
				Message: "User ditemukan",
				Data:    user,
			})
			return
		}
	}
	ctx.JSON(400, models.Response{
		Success: false,
		Message: "User tidak ditemukan",
		Data:    nil,
	})
}

func(u *User) AddUser(ctx *gin.Context) {
	var newUser models.User
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(400, models.Response{
			Success: false,
			Message: "Gagal membaca JSON",
			Data:    nil,
		})
		return
	}

	newUser.Id = len(users) + 1
	users = append(users, newUser)

	ctx.JSON(200, models.Response{
		Success: true,
		Message: "Berhasil menambahkan user",
		Data:    users,
	})
}

func(u *User) EditUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	var updateUser models.User
	if err := ctx.ShouldBindJSON(&updateUser); err != nil {
		ctx.JSON(400, models.Response{
			Success: false,
			Message: "Gagal membaca JSON",
			Data:    nil,
		})
		return
	}

	for i, user := range users {
		if user.Id == id {

			if updateUser.Name != "" {
				users[i].Name = updateUser.Name
			}

			if updateUser.Email != "" {
				users[i].Email = updateUser.Email
			}

			if updateUser.Password != "" {
				argon := argon2.DefaultConfig()
				encoded, err := argon.HashEncoded([]byte(updateUser.Password))
				if err != nil {
					ctx.JSON(400, models.Response{
						Success: false,
						Message: "Gagal hash password baru",
						Data:    nil,
					})
					return
				}
				users[i].Password = string(encoded)
			}

			ctx.JSON(200, models.Response{
				Success: true,
				Message: "Data berhasil diperbarui",
				Data:    users[i],
			})
			return
		}
	}
	ctx.JSON(400, models.Response{
		Success: false,
		Message: "User Tidak Di temukan",
	})
}

func(u *User) DeleteUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)
	for i, user := range users {
		if user.Id == id {
			users = append(users[:i], users[i+1:]...)

			ctx.JSON(200, models.Response{
				Success: true,
				Message: "User Behasil di hapus",
				Data:    users,
			})
			return
		}
	}
	ctx.JSON(400, models.Response{
		Success: false,
		Message: "User Tidak Ditemukan",
	})

}
