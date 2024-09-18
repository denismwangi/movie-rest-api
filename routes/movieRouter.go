package movieRouter

import (
	"github.com/gin-gonic/gin"
	"movie-rest-api/controllers"
)

func MoviesRoutes(router *gin.Engine) {
	moviesGroup := router.Group("/movies")
	{
		moviesGroup.GET("/", controllers.GetAllMovies)
		moviesGroup.POST("/create", controllers.CreateMovieHandler)
		moviesGroup.GET("/:id", controllers.GetMovieById)
	}
}
