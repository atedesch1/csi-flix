package database

func (db *Database) GetAllMovies() ([]Movie, error) {
	var movies []Movie

	tx := db.Db.
		Preload("Directors").
		Preload("Cast").
		Preload("Countries").
		Preload("Genres").
		Find(&movies)

	return movies, tx.Error
}

func (db *Database) GetMovieById(id string) (Movie, error) {
	var movie Movie

	tx := db.Db.
		Preload("Directors").
		Preload("Cast").
		Preload("Countries").
		Preload("Genres").
		First(&movie, id)

	return movie, tx.Error
}

func (db *Database) GetMoviesByTitle(title string) ([]Movie, error) {
	var movies []Movie

	tx := db.Db.
		Preload("Directors").
		Preload("Cast").
		Preload("Countries").
		Preload("Genres").
		Where("Title iLIKE ?", "%"+title+"%").
		Find(&movies)

	return movies, tx.Error
}
