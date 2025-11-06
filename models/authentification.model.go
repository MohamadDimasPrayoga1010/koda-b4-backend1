package models

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type User struct {
	Id       int    `form:"id"`
	Name     string `form:"name"`
	Email    string `form:"email" json:""`
	Password string `form:"password"`
	ProfilePicture string `form:"picture"`
}

var Users = []User{
	{Id: 1, Name: "Yoga", Email: "y@mail.com", Password: "123"},
	{Id: 2, Name: "Dimas", Email: "d@mail.com", Password: "1"},
}

type RegisterRequest struct {
	Name     string `form:"name" example:"Yoga"`
	Email    string `form:"email" example:"yoga@mail.com"`
	Password string `form:"password" example:"123456"`
}

type LoginRequest struct {
	Email    string `form:"email" example:"yoga@mail.com"`
	Password string `form:"password" example:"123456"`
}