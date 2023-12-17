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
	songs, err := repository.GetAllCollections()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving song : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return songs, nil
}

func GetCollectionById(id uuid.UUID) (*models.Song, error) {
	song, err := repository.GetCollectionById(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.CustomError{
				Message: "song not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving song : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return song, err
}

func PutCollectionById(id uuid.UUID, songName string, songAuthor string, songGenre string) (*models.Song, error) {
	_, err := repository.PutCollectionById(id, songName, songAuthor, songGenre)
	if err != nil {

		logrus.Errorf("error updating song : %s", err.Error())
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

		logrus.Errorf("error updating song : %s", err.Error())
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

		logrus.Errorf("error deleting song : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}
	return nil, err
}
