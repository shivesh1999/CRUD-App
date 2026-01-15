package repository

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	"github.com/shivesh/crud-app/database/migrations"
	"github.com/shivesh/crud-app/database/models"
	"gopkg.in/go-playground/validator.v9"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateStruct(contact models.Contact) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(contact)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func (r *Repository) CreateContact(c *gin.Context) {
	var contact models.Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Request failed"})
		return
	}
	errors := ValidateStruct(contact)
	if errors != nil {
		c.JSON(http.StatusBadRequest, errors)
		return
	}
	if err := r.DB.Create(&contact).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Couldn't create contact", "data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Contact has been added", "data": contact})
}
func (r *Repository) UpdateContact(c *gin.Context) {
	var contact models.Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Request failed"})
		return
	}
	errors := ValidateStruct(contact)
	if errors != nil {
		c.JSON(http.StatusBadRequest, errors)
		return
	}
	db := r.DB
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "ID cannot be empty"})
		return
	}
	if db.Model(&contact).Where("id = ?", id).Updates(&contact).RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not get User with given id"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User successfully updated"})
}

func (r *Repository) DeleteContact(c *gin.Context) {
	contactModel := migrations.Contacts{}
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "ID cannot be empty"})
		return
	}
	err := r.DB.Delete(&contactModel, id)
	if err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not delete contact"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Contact delete successfully"})
}
func (r *Repository) GetContacts(c *gin.Context) {
	db := r.DB
	model := db.Model(&migrations.Contacts{})

	pg := paginate.New(&paginate.Config{
		DefaultSize:        20,
		CustomParamEnabled: true,
	})

	page := pg.With(model).Request(c.Request).Response(&[]migrations.Contacts{})

	c.JSON(http.StatusOK, gin.H{"data": page})
}

func (r *Repository) GetContactByID(c *gin.Context) {
	id := c.Param("id")
	contactModel := &migrations.Contacts{}
	if id == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "ID cannot be empty"})
		return
	}
	err := r.DB.Where("id = ?", id).First(contactModel).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not get the contact"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Contact id fetched successfully", "data": contactModel})
}
