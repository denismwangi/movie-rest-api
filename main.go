package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"movie-rest-api/config"
	"movie-rest-api/routes"
)

func main() {
	db, err := connectDb.DbConnection()
	if err != nil {
		fmt.Println("Database connection failed:", err)
		return
	}
	defer db.Close()

	fmt.Println("Database connected successfully!")

	// Initialize Gin router
	router := gin.Default()

	movieRouter.MoviesRoutes(router)

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to the Movie API")
	})

	router.Run(":8080")
}
