package main

import (
	"net/http"

	"github.com/atedesch1/csi-flix/cmd/database"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.SetupDB("postgresql://root:root@localhost:5432/core")
	db.Migrate()
	movies := database.ReadMovieCsv("assets/netflix_titles.csv")
	db.PopulateDB(movies)

	r := gin.Default()

	r.GET("/movies", func(c *gin.Context) {
		var movies []database.Movie

		db.Db.Preload("Directors").Preload("Cast").Preload("Countries").Preload("Genres").Find(&movies)

		c.JSON(http.StatusOK, movies)
	})

	r.Run()
}
