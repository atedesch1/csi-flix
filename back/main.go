package main

import (
	"github.com/atedesch1/csi-flix/cmd/database"
)

func main() {
	db := database.SetupDB("postgresql://root:root@localhost:5432/core")
	movies := database.ReadMovieCsv("assets/netflix_titles.csv")
	db.PopulateDB(movies)
}
