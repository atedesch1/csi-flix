package db

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model `json:"-"`

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

	Reviews []Review `gorm:"foreignKey:MovieID"`
}

type Director struct {
	gorm.Model `json:"-"`

	Name string `gorm:"index:,unique"`
}

type Actor struct {
	gorm.Model `json:"-"`

	Name string `gorm:"index:,unique"`
}

type Country struct {
	gorm.Model `json:"-"`

	Name string `gorm:"index:,unique"`
}

type Genre struct {
	gorm.Model `json:"-"`

	Name string `gorm:"index:,unique"`
}

type Review struct {
	gorm.Model `json:"-"`
	MovieID    uint

	Rating    uint
	Content   string
	Anonymous bool
	Username  string
}
