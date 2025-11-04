package models

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type User struct {
	Id       int    `form:"id"`
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

var Users = []User{
	{Id: 1, Name: "Yoga"},
	{Id: 2, Name: "Dimas"},
}