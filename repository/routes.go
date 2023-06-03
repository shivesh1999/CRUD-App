package repository

import "github.com/gofiber/fiber/v2"

func (repo *Repository) SetupRoutes(app *fiber.App) {
	//routes.Setup(repo, app)
	app.Static("/", "./client/public")

	api := app.Group("/api")
	api.Get("/contacts", repo.GetContacts)
	api.Post("/contacts", repo.CreateContact)
	api.Patch("/contacts/:id", repo.UpdateContact)
	api.Delete("/contacts/:id", repo.DeleteContact)
	api.Get("/contacts/:id", repo.GetContactByID)
	api.Post("/register", repo.RegisterUser)
	api.Post("/login", repo.LoginUser)
	api.Get("/user", repo.User)
	api.Post("/logout", repo.Logout)
}
