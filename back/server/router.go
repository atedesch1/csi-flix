package server

import (
	"github.com/atedesch1/csi-flix/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	movieController := new(controllers.MovieController)
	reviewController := new(controllers.ReviewController)

	v1 := router.Group("v1")
	{
		movieGroup := v1.Group("movie")
		{
			movieGroup.GET("/", movieController.GetAll)
			movieGroup.GET("/bytitle/:title", movieController.GetByTitle)
			movieGroup.GET("/:id", movieController.GetById)
			movieGroup.GET("/:id/review", reviewController.GetMovieReviews)
			movieGroup.POST("/:id/review", reviewController.CreateReview)
		}

		reviewGroup := v1.Group("review")
		{
			reviewGroup.GET("/", reviewController.GetAll)
			reviewGroup.GET("/:id", reviewController.GetById)
			reviewGroup.DELETE("/:id", reviewController.DeleteReview)
		}
	}

	return router
}
