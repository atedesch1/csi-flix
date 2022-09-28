package database

import (
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

func (db *Database) PopulateDB(movies []*Movie) (tx *gorm.DB) {
	return db.gormDb.CreateInBatches(movies, 1000)
}
