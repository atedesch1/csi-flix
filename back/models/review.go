package models

import (
	"github.com/atedesch1/csi-flix/db"
	"gorm.io/gorm"
)

type Review struct {
	Rating    uint
	Content   string
	Anonymous bool
	Username  string
}

func DbReviewToReview(dbReview db.Review) Review {
	return Review{
		Rating:    dbReview.Rating,
		Content:   dbReview.Content,
		Anonymous: dbReview.Anonymous,
		Username:  dbReview.Username,
	}
}

func DbReviewsToReviews(dbReviews []db.Review) []Review {
	Reviews := make([]Review, len(dbReviews))
	for i, dbReview := range dbReviews {
		Reviews[i] = DbReviewToReview(dbReview)
	}
	return Reviews
}

func (m Review) CreateReview(dbMovie db.Movie, review Review) (db.Review, error) {
	dbReview := &db.Review{
		Rating:    review.Rating,
		Content:   review.Content,
		Anonymous: review.Anonymous,
		Username:  review.Username,
	}

	tx := db.GetDB().Model(&dbMovie).Association("Reviews").Append(dbReview)

	return *dbReview, tx
}

func (m Review) DeleteReview(id uint) error {
	dbReview := db.Review{
		Model: gorm.Model{
			ID: id,
		},
	}

	tx := db.GetDB().Delete(dbReview)

	return tx.Error
}

func (m Review) GetAll(limit int) ([]db.Review, error) {
	var dbReviews []db.Review

	tx := db.GetDB().
		Limit(limit).
		Find(&dbReviews)

	return dbReviews, tx.Error
}

func (m Review) GetById(id string) (db.Review, error) {
	var dbReview db.Review

	tx := db.GetDB().
		First(&dbReview, id)

	return dbReview, tx.Error
}

func (m Review) GetMovieReviews(dbMovie db.Movie) ([]db.Review, error) {
	var dbReviews []db.Review

	err := db.GetDB().Model(&dbMovie).Association("Reviews").Find(&dbReviews)

	return dbReviews, err
}
