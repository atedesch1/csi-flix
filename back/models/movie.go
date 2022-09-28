package models

import (
	"github.com/atedesch1/csi-flix/db"
	"gorm.io/gorm"
)

type Movie struct {
	Type        string
	Title       string
	Directors   []string
	Cast        []string
	Countries   []string
	DateAdded   string
	ReleaseYear string
	Rating      string
	Duration    string
	Genres      []string
	Description string
}

func DbMovieToMovie(dbMovie db.Movie) Movie {
	directors := make([]string, len(dbMovie.Directors))
	for i, director := range dbMovie.Directors {
		directors[i] = director.Name
	}

	cast := make([]string, len(dbMovie.Cast))
	for i, actor := range dbMovie.Cast {
		cast[i] = actor.Name
	}

	countries := make([]string, len(dbMovie.Countries))
	for i, country := range dbMovie.Countries {
		countries[i] = country.Name
	}

	genres := make([]string, len(dbMovie.Genres))
	for i, genre := range dbMovie.Genres {
		genres[i] = genre.Name
	}

	return Movie{
		Type:        dbMovie.Type,
		Title:       dbMovie.Title,
		Directors:   directors,
		Cast:        cast,
		Countries:   countries,
		DateAdded:   dbMovie.DateAdded,
		ReleaseYear: dbMovie.ReleaseYear,
		Rating:      dbMovie.Rating,
		Duration:    dbMovie.Duration,
		Genres:      genres,
		Description: dbMovie.Description,
	}
}

func DbMoviesToMovies(dbMovies []db.Movie) []Movie {
	movies := make([]Movie, len(dbMovies))
	for i, dbMovie := range dbMovies {
		movies[i] = DbMovieToMovie(dbMovie)
	}
	return movies
}

func (m Movie) GetDbWithPreloadedMovieAssociations() *gorm.DB {
	return db.GetDB().
		Preload("Directors").
		Preload("Cast").
		Preload("Countries").
		Preload("Genres")
}

func (m Movie) GetAll() ([]Movie, error) {
	var dbMovies []db.Movie

	tx := m.GetDbWithPreloadedMovieAssociations().
		Find(&dbMovies)

	return DbMoviesToMovies(dbMovies), tx.Error
}

func (m Movie) GetById(id string) (Movie, error) {
	var dbMovie db.Movie

	tx := m.GetDbWithPreloadedMovieAssociations().
		First(&dbMovie, id)

	return DbMovieToMovie(dbMovie), tx.Error
}

func (m Movie) GetByTitle(title string) ([]Movie, error) {
	var dbMovies []db.Movie

	tx := m.GetDbWithPreloadedMovieAssociations().
		Where("Title iLIKE ?", "%"+title+"%").
		Find(&dbMovies)

	return DbMoviesToMovies(dbMovies), tx.Error
}
