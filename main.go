package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shivesh/crud-app/bootstrap"
)

func main() {
	router := gin.Default()
	bootstrap.InitializeApp(router)
}
