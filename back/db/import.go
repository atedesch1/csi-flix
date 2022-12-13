package db

import (
	"os"

	"github.com/atedesch1/csi-flix/utils"
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

func ReadMovieCsv(path string) []*Movie {
	csv, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		panic(err)
	}
	defer csv.Close()

	moviesCsv := []*MovieCsv{}

	if err := gocsv.UnmarshalFile(csv, &moviesCsv); err != nil {
		panic(err)
	}

	movies := make([]*Movie, len(moviesCsv))

	for idx, movieCsv := range moviesCsv {
		movies[idx] = movieCsv.ToMovie()
	}

	return movies
}

func (m *MovieCsv) ToMovie() *Movie {
	directorNames := utils.SplitAndTrim(m.Directors, ",")
	directors := make([]Director, len(directorNames))
	for i, directorName := range directorNames {
		directors[i] = Director{Name: directorName}
	}

	actorNames := utils.SplitAndTrim(m.Cast, ",")
	cast := make([]Actor, len(actorNames))
	for i, actorName := range actorNames {
		cast[i] = Actor{Name: actorName}
	}

	countryNames := utils.SplitAndTrim(m.Countries, ",")
	countries := make([]Country, len(countryNames))
	for i, countryName := range countryNames {
		countries[i] = Country{Name: countryName}
	}

	genreNames := utils.SplitAndTrim(m.Genres, ",")
	genres := make([]Genre, len(genreNames))
	for i, genreName := range genreNames {
		genres[i] = Genre{Name: genreName}
	}

	return &Movie{
		Type:        m.Type,
		Title:       m.Title,
		Directors:   directors,
		Cast:        cast,
		Countries:   countries,
		DateAdded:   m.DateAdded,
		ReleaseYear: m.ReleaseYear,
		Rating:      m.Rating,
		Duration:    m.Duration,
		Genres:      genres,
		Description: m.Description,
		Reviews:     make([]Review, 0),
	}
}
