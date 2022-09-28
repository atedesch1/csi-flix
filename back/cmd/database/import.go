package database

import (
	"os"

	"github.com/gocarina/gocsv"
)

type MovieCsv struct {
	Id          string `csv:"show_id"`
	Type        string `csv:"type"`
	Title       string `csv:"title"`
	Directors   string `csv:"director"`
	Cast        string `csv:"cast"`
	Countries   string `csv:"country"`
	DateAdded   string `csv:"date_added"`
	ReleaseYear string `csv:"release_year"`
	Rating      string `csv:"rating"`
	Duration    string `csv:"duration"`
	Genres      string `csv:"listed_in"`
	Description string `csv:"description"`
}

func ReadMovieCsv(path string) []*MovieCsv {
	csv, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		panic(err)
	}
	defer csv.Close()

	movies := []*MovieCsv{}

	if err := gocsv.UnmarshalFile(csv, &movies); err != nil {
		panic(err)
	}

	return movies
}
