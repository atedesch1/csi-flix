package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	var err error
	db, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Movie{})
	db.AutoMigrate(&Review{})
	movies := ReadMovieCsv("assets/netflix_titles.csv")
	db.CreateInBatches(movies, 1000)
}

func GetDB() *gorm.DB {
	return db
}
