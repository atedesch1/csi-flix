package main

import (
	"net/http"
	"os"

	"github.com/atedesch1/csi-flix/cmd/database"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.SetupDB(os.Getenv("DATABASE_URL"))
	db.Migrate()
	movies := database.ReadMovieCsv("assets/netflix_titles.csv")
	db.PopulateDB(movies)

	r := gin.Default()

	r.GET("/movies", func(c *gin.Context) {
		movies, err := db.GetAllMovies()

		if err != nil {
			c.String(http.StatusInternalServerError, "Internal Server Error")
		} else {
			c.JSON(http.StatusOK, movies)
		}
	})

	r.GET("/movies/:id", func(c *gin.Context) {
		id := c.Param("id")

		movie, err := db.GetMovieById(id)

		if err != nil {
			c.JSON(http.StatusNotFound, nil)
		} else {
			c.JSON(http.StatusOK, movie)
		}
	})

	r.GET("/movies/bytitle/:title", func(c *gin.Context) {
		title := c.Param("title")

		movies, err := db.GetMoviesByTitle(title)

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
		} else {
			c.JSON(http.StatusOK, movies)
		}
	})

	r.Run()
}
