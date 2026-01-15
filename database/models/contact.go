package models

type Contact struct {
	Name         string `json:"name" validate:"required,min=5,max=40"`
	MobileNumber string `json:"MobileNumber" validate:"required,min=10,max=10"`
	City         string `json:"city" validate:"required,max=40"`
	Country      string `json:"country" validate:"required,max=40"`
	Email        string `json:"email" validate:"required,email,min=10,max=40"`
}
