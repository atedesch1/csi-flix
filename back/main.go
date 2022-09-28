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
		var movies []database.Movie

		db.Db.
			Preload("Directors").
			Preload("Cast").
			Preload("Countries").
			Preload("Genres").
			Find(&movies)

		c.JSON(http.StatusOK, movies)
	})

	r.GET("/movies/:id", func(c *gin.Context) {
		var movie database.Movie
		id := c.Param("id")

		tx := db.Db.
			Preload("Directors").
			Preload("Cast").
			Preload("Countries").
			Preload("Genres").
			First(&movie, id)

		if tx.Error != nil {
			c.JSON(http.StatusNotFound, nil)
		} else {
			c.JSON(http.StatusOK, movie)
		}
	})

	r.GET("/movies/bytitle/:title", func(c *gin.Context) {
		var movies []database.Movie
		title := c.Param("title")

		tx := db.Db.
			Preload("Directors").
			Preload("Cast").
			Preload("Countries").
			Preload("Genres").
			Where("Title iLIKE ?", "%"+title+"%").
			Find(&movies)

		if tx.Error != nil {
			c.JSON(http.StatusInternalServerError, nil)
		} else {
			c.JSON(http.StatusOK, movies)
		}
	})

	r.Run()
}
