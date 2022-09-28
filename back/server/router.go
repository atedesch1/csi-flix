package server

import (
	"github.com/atedesch1/csi-flix/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("v1")
	{
		movieGroup := v1.Group("movie")
		{
			controller := new(controllers.MovieController)
			movieGroup.GET("/", controller.GetAll)
			movieGroup.GET("/:id", controller.GetById)
			movieGroup.GET("/bytitle/:title", controller.GetByTitle)
		}
	}

	return router
}
