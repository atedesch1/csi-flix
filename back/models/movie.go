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
	Reviews     []Review
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

	reviews := make([]Review, 0)
	for _, dbReview := range dbMovie.Reviews {
		reviews = append(reviews, DbReviewToReview(dbReview))
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
		Reviews:     reviews,
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
		Preload("Genres").
		Preload("Reviews")

}

func (m Movie) GetAll(limit int) ([]db.Movie, error) {
	var dbMovies []db.Movie

	tx := m.GetDbWithPreloadedMovieAssociations().
		Limit(limit).
		Find(&dbMovies)

	return dbMovies, tx.Error
}

func (m Movie) GetById(id string) (db.Movie, error) {
	var dbMovie db.Movie

	tx := m.GetDbWithPreloadedMovieAssociations().
		First(&dbMovie, id)

	return dbMovie, tx.Error
}

func (m Movie) GetByTitle(title string) ([]db.Movie, error) {
	var dbMovies []db.Movie

	tx := m.GetDbWithPreloadedMovieAssociations().
		Where("Title iLIKE ?", "%"+title+"%").
		Find(&dbMovies)

	return dbMovies, tx.Error
}
