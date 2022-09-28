package main

import (
	"github.com/atedesch1/csi-flix/cmd/database"
)

func main() {
	csv := database.ReadMovieCsv("assets/netflix_titles.csv")
	db := database.SetupDB("postgresql://root:root@localhost:5432/core")
	db.PopulateDB(csv)
}
