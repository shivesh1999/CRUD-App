package bootstrap

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/shivesh/crud-app/database/migrations"
	"github.com/shivesh/crud-app/database/storage"
	"github.com/shivesh/crud-app/repository"
)

func InitializeApp(router *gin.Engine) {
	_, ok := os.LookupEnv("APP_ENV")

	if !ok {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal(err)
		}
	}
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("could not load the database")
	}

	err = migrations.MigrateContacts(db)
	if err != nil {
		log.Fatal("Could not migrate db")
	}

	repo := repository.Repository{
		DB: db,
	}

	router.Use(cors.Default())
	repo.SetupRoutes(router)
	router.Run(":8080")
}
