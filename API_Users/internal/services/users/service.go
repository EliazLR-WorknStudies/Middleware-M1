package users

import (
	"database/sql"
	"errors"
	"middleware/users/internal/models"
	repository "middleware/users/internal/repositories/users"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func GetAllUsers() ([]models.User, error) {
	var err error
	// calling repository
	users, err := repository.GetAllUsers()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving users : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return users, nil
}

func GetUserById(id uuid.UUID) (*models.User, error) {
	user, err := repository.GetUserById(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "user not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving users : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return user, err
}

func CreateUserByUsername(user *models.User) (*models.User, error) {
	id := uuid.Must(uuid.NewV4())
	user.Id = &id
	err := repository.CreateUser(user)
	if err != nil {
		logrus.Errorf("error creating user : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return user, nil
}

func DeleteUserById(id uuid.UUID) error {
	err := repository.DeleteUser(id)
	if err != nil {
		logrus.Errorf("error deleting user : %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return nil
}

func UpdateUser(user *models.User) (*models.User, error) {
	err := repository.UpdateUser(user)
	if err != nil {
		logrus.Errorf("error updating user : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return user, nil
}
