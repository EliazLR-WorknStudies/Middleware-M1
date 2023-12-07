package ratings

import (
	"database/sql"
	"middleware/ratings/internal/models"
	repository "middleware/ratings/internal/repositories/ratings"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
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
		if err.Error() == sql.ErrNoRows.Error() {
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
	id, err := uuid.NewV4()
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

func UpdateRating(rating *models.Ratings) (*models.Ratings, error) {

	collection, err := repository.UpdateRating(rating)

	if err != nil {
		logrus.Errorf("error adding rating : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return collection, err
}

func DeleteRating(id uuid.UUID) (*models.Ratings, error) {
	collection, err := repository.DeleteRating(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
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

func GetRatingBySongId(id uuid.UUID) ([]models.Ratings, error) {
	collection, err := repository.GetRatingBySongId(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
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
