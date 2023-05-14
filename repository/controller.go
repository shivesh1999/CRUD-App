package repository

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
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

func (r *Repository) CreateContact(context *fiber.Ctx) error {
	contact := models.Contact{}
	err := context.BodyParser(&contact)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})

		return err
	}
	errors := ValidateStruct(contact)
	if errors != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errors)
	}

	if err := r.DB.Create(&contact).Error; err != nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Couldn't create contact", "data": err})
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Contact has been added", "data": contact})
	return nil
}
func (r *Repository) UpdateContact(context *fiber.Ctx) error {
	contact := models.Contact{}
	err := context.BodyParser(&contact)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})

		return err
	}
	errors := ValidateStruct(contact)
	if errors != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errors)
	}

	db := r.DB
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}
	if db.Model(&contact).Where("id = ?", id).Updates(&contact).RowsAffected == 0 {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Could not get User with given id"})
	}

	return context.JSON(fiber.Map{"status": "success", "message": "User successfully updated"})
}

func (r *Repository) DeleteContact(context *fiber.Ctx) error {
	contactModel := migrations.Contacts{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}

	err := r.DB.Delete(contactModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not delete contact"})
		return err.Error
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Contact delete successfully"})
	return nil
}
func (r *Repository) GetContacts(context *fiber.Ctx) error {
	db := r.DB
	model := db.Model(&migrations.Contacts{})

	pg := paginate.New(&paginate.Config{
		DefaultSize:        20,
		CustomParamEnabled: true,
	})

	page := pg.With(model).Request(context.Request()).Response(&[]migrations.Contacts{})

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"data": page,
	})
	return nil
}

func (r *Repository) GetContactByID(context *fiber.Ctx) error {
	id := context.Params("id")
	contactModel := &migrations.Contacts{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(contactModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Could not get the contact"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Contact id fetched successfully", "data": contactModel})
	return nil
}
