package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Url string
	Db  *gorm.DB
}

func SetupDB(url string) *Database {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database := &Database{
		Url: url,
		Db:  db,
	}

	return database
}

func (db *Database) Migrate() error {
	return db.Db.AutoMigrate(&Movie{})
}

func (db *Database) PopulateDB(movies []*Movie) (tx *gorm.DB) {
	return db.Db.CreateInBatches(movies, 1000)
}
