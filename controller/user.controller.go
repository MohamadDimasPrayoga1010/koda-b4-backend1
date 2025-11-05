package controller

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
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
// GetUsers godoc
// @Summary Ambil daftar user dengan paginasi dan pencarian
// @Description Menampilkan daftar user berdasarkan halaman (page), limit, dan kata kunci pencarian (search)
// @Tags Users
// @Produce json
// @Param page query int false "Nomor halaman (default: 1)"
// @Param limit query int false "Jumlah data per halaman (default: 10)"
// @Param search query string false "Kata kunci pencarian nama user"
// @Success 200 {object} models.Response
// @Router /users [get]
func (u *User) GetUsers(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "10")
	search := strings.ToLower(ctx.DefaultQuery("search", ""))

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	filteredUsers := []models.User{}
	for _, user := range users {
		if search == "" || strings.Contains(strings.ToLower(user.Name), search) {
			filteredUsers = append(filteredUsers, user)
		}
	}

	totalData := len(filteredUsers)
	start := (page - 1) * limit
	end := start + limit

	if start > totalData {
		start = totalData
	}
	if end > totalData {
		end = totalData
	}

	pagedUsers := filteredUsers[start:end]

	totalPages := (totalData + limit - 1) / limit 

	ctx.JSON(200, models.Response{
		Success: true,
		Message: "Berhasil ambil data user",
		Data: gin.H{
			"page":         page,
			"limit":        limit,
			"search":       search,
			"total_data":   totalData,
			"total_pages":  totalPages,
			"users":        pagedUsers,
		},
	})
}

// GetUserId godoc
// @Summary Ambil user berdasarkan ID
// @Description Mendapatkan data user sesuai ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "ID User"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /users/{id} [get]
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

// AddUser godoc
// @Summary Tambah user baru
// @Description Menambahkan user baru ke sistem
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "Data user baru"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /users [post]
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

// EditUser godoc
// @Summary Edit data user
// @Description Mengedit data user berdasarkan ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "ID User"
// @Param user body models.User true "Data user yang diupdate"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /users/{id} [patch]
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

// DeleteUser godoc
// @Summary Hapus user
// @Description Menghapus user berdasarkan ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "ID User"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /users/{id} [delete]
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
