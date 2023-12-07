package users

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"middleware/users/internal/models"
	"middleware/users/internal/services/users"
	"net/http"
)

// CreateUser
// @Tags         users
// @Summary      Creates a user with username.
// @Description  Creates a user with username.
// @Param        username      body      string  true  "Username"
// @Success      200            {object}  models.User
// @Failure      422            "Cannot parse username"
// @Failure      500            "Something went wrong"
// @Router       /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// parsing body
	var user *models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logrus.Errorf("parsing error : %s", err.Error())
		customError := &models.CustomError{
			Message: "Cannot parse username",
			Code:    http.StatusUnprocessableEntity,
		}
		w.WriteHeader(customError.Code)
		body, _ := json.Marshal(customError)
		_, _ = w.Write(body)
		return
	}

	// calling service
	user, err = users.CreateUserByUsername(user)
	if err != nil {
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(user)
	_, _ = w.Write(body)
	return
}