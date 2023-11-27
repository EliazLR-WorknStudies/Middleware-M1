package raitings

import (
	"database/sql"
	"errors"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/ratings/internal/models"
	repository "middleware/ratings/internal/repositories/ratings"
	"net/http"
)

func GetAllRatings() ([]models.Ratings, error) {
	var err error
	// calling repository
	collections, err := repository.GetAllRatings()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return collections, nil
}

func GetRatingById(id uuid.UUID) (*models.Ratings, error) {
	collection, err := repository.GetRatingById(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "collection not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return collection, err
}


func CreateRating(rating *models.Ratings) (*models.Ratings, error) {
	id, err:= uuid.NewV4()
	rating.Id = &id
	
	collection, err := repository.CreateRating(rating)
	
	if err != nil {
		logrus.Errorf("error adding rating : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return collection, err
}
