package controllers

import (
	"net/http"
	"strconv"

	"github.com/atedesch1/csi-flix/api"
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

	dbMovies, err := movieModel.GetAll(limit)
	movies := models.DbMoviesToMovies(dbMovies)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't retrieve movies", "error": err})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, movies)
}

func (u MovieController) GetById(c *gin.Context) {
	if c.Param("id") != "" {
		dbMovie, err := movieModel.GetById(c.Param("id"))
		movie := models.DbMovieToMovie(dbMovie)

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
	stats := c.Query("stats")
	title := c.Param("title")

	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		return
	}

	dbMovies, err := movieModel.GetByTitle(title)
	movies := models.DbMoviesToMovies(dbMovies)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't retrieve movies", "error": err})
		c.Abort()
		return
	}

	if stats == "true" {
		result := make([]map[string]interface{}, len(movies))

		for i, movie := range movies {
			stats := api.GetMovieStatsFromApi(movie.Title)
			result[i] = map[string]interface{}{
				"Movie": movie,
				"Stats": stats,
			}
		}
		c.JSON(http.StatusOK, result)
		return
	}

	c.JSON(http.StatusOK, movies)
}
