package repository

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
	"github.com/shivesh/crud-app/database/migrations"
	"github.com/shivesh/crud-app/database/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

const secretKey = "secret"

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

func (r *Repository) RegisterUser(context *fiber.Ctx) error {
	data := models.User{}
	err := context.BodyParser(&data)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})

		return err
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	user := models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: string(password),
	}
	if err := r.DB.Create(&user).Error; err != nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "User has been added", "data": user})
	return nil
}

func (r *Repository) LoginUser(context *fiber.Ctx) error {
	data := models.LoginData{}
	err := context.BodyParser(&data)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})

		return err
	}

	var user models.User

	r.DB.Where("email = ?", data.Email).First(&user)
	if user.ID == 0 {
		context.Status(fiber.StatusNotFound)
		return context.JSON(fiber.Map{
			"message": "User not found",
		})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		context.Status(fiber.StatusBadRequest)
		return context.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		context.Status(fiber.StatusInternalServerError)
		return context.JSON(fiber.Map{
			"message": "Could not log in",
		})
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	context.Cookie(&cookie)
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "User logged in successfully"})
	return nil
}

func (r *Repository) User(context *fiber.Ctx) error {
	cookie := context.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		context.Status(fiber.StatusUnauthorized)
		return context.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)
	var user models.User
	r.DB.Where("id = ?", claims.Issuer).First(&user)

	return context.JSON(user)
}

func (r *Repository) Logout(context *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	context.Cookie(&cookie)
	return context.JSON(fiber.Map{
		"message": "User logged out successfully",
	})
}
