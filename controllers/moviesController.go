package controllers

import (
	"movie-rest-api/services"
	"movie-rest-api/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetAllMovies(c *gin.Context) {
	movieItems, err := services.ListMoviesHandler()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(movieItems) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No movies found"})
	} else {
		c.JSON(http.StatusOK, movieItems)
	}
}

func CreateMovieHandler(c *gin.Context) {
	var movieItem models.Movie
	if err := c.BindJSON(&movieItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := services.CreateMovieHandler(movieItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, movieItem)
}

func GetMovieById(c *gin.Context) {
	id := c.Param("id")
	movieItem, err := services.GetMovieById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if movieItem == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Movie not found"})
	} else {
		c.IndentedJSON(http.StatusOK, movieItem)
	}
}
