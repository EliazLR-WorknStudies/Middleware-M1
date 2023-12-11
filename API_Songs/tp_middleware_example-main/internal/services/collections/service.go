package collections

import (
	"database/sql"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/collections"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func GetAllCollections() ([]models.Song, error) {
	var err error
	// calling repository
	collections, err := repository.GetAllCollections()
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

func GetCollectionById(id uuid.UUID) (*models.Song, error) {
	collection, err := repository.GetCollectionById(id)
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

func PutCollectionById(id uuid.UUID, songName string, songAuthor string, songGenre string) (*models.Song, error) {
	_, err := repository.PutCollectionById(id, songName, songAuthor, songGenre)
	if err != nil {

		logrus.Errorf("error updating collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}
	return nil, err
}

func PostCollectionById(id uuid.UUID, songName string, songAuthor string, songGenre string) (*models.Song, error) {
	_, err := repository.CreateCollectionByRepo(id, songName, songAuthor, songGenre)
	if err != nil {

		logrus.Errorf("error updating collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}
	return nil, err
}

func DeleteCollectionById(id uuid.UUID) (*models.Song, error) {
	_, err := repository.DeleteCollectionByRepo(id)
	if err != nil {

		logrus.Errorf("error deleting collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}
	return nil, err
}
