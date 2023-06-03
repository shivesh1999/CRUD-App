package models

type LoginData struct {
	Email    string `json : "email" validate:"required,email,min=10,max=40"`
	Password string `json : "password" validate:"required,min=8,max=28"`
}
