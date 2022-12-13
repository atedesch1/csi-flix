package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/atedesch1/csi-flix/models"
	"github.com/gin-gonic/gin"
)

type ReviewController struct{}

var reviewModel = new(models.Review)

func (u ReviewController) GetAll(c *gin.Context) {
	limitQuery := c.Query("limit")

	if limitQuery == "" {
		limitQuery = "100"
	}

	limit, err := strconv.Atoi(limitQuery)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "limit query badly formatted"})
		c.Abort()
		return
	}

	dbReviews, err := reviewModel.GetAll(limit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't retrieve reviews", "error": err})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, models.DbReviewsToReviews(dbReviews))
}

func (u ReviewController) GetById(c *gin.Context) {
	if c.Param("id") != "" {
		dbReview, err := reviewModel.GetById(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't retrieve review", "error": err})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, models.DbReviewToReview(dbReview))
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
}

func (u ReviewController) GetMovieReviews(c *gin.Context) {
	if c.Param("id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		return
	}

	dbMovie, err := movieModel.GetById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't retrieve movie", "error": err})
		c.Abort()
		return
	}

	dbReviews, err := reviewModel.GetMovieReviews(dbMovie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't retrieve movie reviews", "error": err})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, models.DbReviewsToReviews(dbReviews))
}

func (u ReviewController) CreateReview(c *gin.Context) {
	if c.Param("id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		return
	}

	dbMovie, err := movieModel.GetById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't retrieve movie", "error": err})
		c.Abort()
		return
	}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't read body", "error": err})
		c.Abort()
		return
	}

	var review models.Review
	err = json.Unmarshal(body, &review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't unmarshal review", "error": err})
		c.Abort()
		return
	}

	dbReview, err := reviewModel.CreateReview(dbMovie, review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't create review", "error": err})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, models.DbReviewToReview(dbReview))
}

func (u ReviewController) DeleteReview(c *gin.Context) {
	if c.Param("id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		return
	}

	dbReview, err := reviewModel.GetById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't retrieve review", "error": err})
		c.Abort()
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	err = reviewModel.DeleteReview(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't delete review", "error": err})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, dbReview)
}
