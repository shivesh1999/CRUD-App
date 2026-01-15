package repository

import "github.com/gin-gonic/gin"

func (repo *Repository) SetupRoutes(router *gin.Engine) {
	router.Static("/static", "./client/public")

	api := router.Group("/api")
	api.GET("/contacts", repo.GetContacts)
	api.POST("/contacts", repo.CreateContact)
	api.PATCH("/contacts/:id", repo.UpdateContact)
	api.DELETE("/contacts/:id", repo.DeleteContact)
	api.GET("/contacts/:id", repo.GetContactByID)
}
