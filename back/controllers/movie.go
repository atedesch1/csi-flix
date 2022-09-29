package controllers

import (
	"net/http"

	"github.com/atedesch1/csi-flix/models"
	"github.com/gin-gonic/gin"
)

type MovieController struct{}

var movieModel = new(models.Movie)

func (u MovieController) GetAll(c *gin.Context) {
	limitQuery := c.Query("limit")

	if limitQuery == "" {
		limitQuery = "100"
	}

	limit, err := strconv.Atoi(limitQuery)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "limit query badly formatted"})
		c.Abort()
		return
	}

	movies, err := movieModel.GetAll(limit)

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
