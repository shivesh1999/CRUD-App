package models

type User struct {
	ID       uint   `json :"id"`
	Name     string `json :"name" validate:"required,min=5,max=40" `
	Email    string `json : "email" validate:"required,email,min=10,max=40"`
	Password string `json :"-",omitempty`
}
