package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model

	Id   uint
	Name string
}

func main() {
	dsn := "postgresql://root:root@localhost:5432/core"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Movie{})

	// Create
	db.Create(&Movie{Id: 1, Name: "Test"})

	// Read
	var movie Movie
	db.First(&movie, 1)                  // find movie with integer primary key
	db.First(&movie, "Name = ?", "Test") // find movie with code D42

	// // Update - update movie's price to 200
	// db.Model(&movie).Update("Price", 200)
	// // Update - update multiple fields
	// db.Model(&movie).Updates(Movie{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&movie).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - delete movie
	// db.Delete(&movie, 1)
}
