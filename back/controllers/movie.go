package controllers

import (
	"net/http"

	"github.com/atedesch1/csi-flix/models"
	"github.com/gin-gonic/gin"
)

type MovieController struct{}

var movieModel = new(models.Movie)

func (u MovieController) GetAll(c *gin.Context) {
	movies, err := movieModel.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't retrieve movies", "error": err})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, movies)
}

func (u MovieController) GetById(c *gin.Context) {
	if c.Param("id") != "" {
		movie, err := movieModel.GetById(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't retrieve movie", "error": err})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, movie)
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
}

func (u MovieController) GetByTitle(c *gin.Context) {
	if c.Param("title") != "" {
		movies, err := movieModel.GetByTitle(c.Param("title"))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't retrieve movies", "error": err})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, movies)
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
}
