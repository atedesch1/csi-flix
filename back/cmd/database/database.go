package database

import (
	"github.com/atedesch1/csi-flix/cmd/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model

	Type        string
	Title       string     `gorm:"index:,unique"`
	Directors   []Director `gorm:"many2many:movie_directors;references:Name"`
	Cast        []Actor    `gorm:"many2many:movie_actors;references:Name"`
	Countries   []Country  `gorm:"many2many:movie_countries;references:Name"`
	DateAdded   string
	ReleaseYear string
	Rating      string
	Duration    string
	Genres      []Genre `gorm:"many2many:movie_genres;references:Name"`
	Description string
}

type Director struct {
	gorm.Model

	Name string `gorm:"index:,unique"`
}

type Actor struct {
	gorm.Model

	Name string `gorm:"index:,unique"`
}

type Country struct {
	gorm.Model

	Name string `gorm:"index:,unique"`
}

type Genre struct {
	gorm.Model

	Name string `gorm:"index:,unique"`
}

type Database struct {
	url    string
	gormDb *gorm.DB
}

func SetupDB(url string) *Database {
	dsn := "postgresql://root:root@localhost:5432/core"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database := &Database{
		url:    url,
		gormDb: db,
	}

	database.gormDb.AutoMigrate(&Movie{})

	return database
}

func (db *Database) PopulateDB(csv []*MovieCsv) (tx *gorm.DB) {
	movies := make([]*Movie, len(csv))

	for idx, movie := range csv {
		directorNames := utils.SplitAndTrim(movie.Directors, ",")
		directors := make([]Director, len(directorNames))
		for i, directorName := range directorNames {
			directors[i] = Director{Name: directorName}
		}

		actorNames := utils.SplitAndTrim(movie.Cast, ",")
		cast := make([]Actor, len(actorNames))
		for i, actorName := range actorNames {
			cast[i] = Actor{Name: actorName}
		}

		countryNames := utils.SplitAndTrim(movie.Countries, ",")
		countries := make([]Country, len(countryNames))
		for i, countryName := range countryNames {
			countries[i] = Country{Name: countryName}
		}

		genreNames := utils.SplitAndTrim(movie.Genres, ",")
		genres := make([]Genre, len(genreNames))
		for i, genreName := range genreNames {
			genres[i] = Genre{Name: genreName}
		}

		movies[idx] = &Movie{
			Type:        movie.Type,
			Title:       movie.Title,
			Directors:   directors,
			Cast:        cast,
			Countries:   countries,
			DateAdded:   movie.DateAdded,
			ReleaseYear: movie.ReleaseYear,
			Rating:      movie.Rating,
			Duration:    movie.Duration,
			Genres:      genres,
			Description: movie.Description,
		}
	}

	return db.gormDb.CreateInBatches(movies, 1000)
}
