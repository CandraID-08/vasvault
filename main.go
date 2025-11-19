package main

import (
	"vasvault/internal/repositories"
	"vasvault/internal/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
   		fmt.Println("No .env file found, using system environment")
	}

	db, err := repositories.Connect()
	if err != nil {
		panic("Failed to connect to database")
	}

	r := gin.Default()
	routes.InitRoutes(r, db.DB)

	if err := r.Run(); err != nil {
		panic("Failed to run server")
	}
}
