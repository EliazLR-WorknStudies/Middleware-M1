package users

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/users/internal/models"
	"middleware/users/internal/services/users"
	"net/http"
)

// UpdateUser
// @Tags         users
// @Summary      Update a user.
// @Description  Update a user.
// @Param        id           	path      string  true  "User UUID formatted ID"
// @Param        username     	body      string  true  "User username"
// @Success      200            {object}  string
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /users/{id} [put]

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId, _ := ctx.Value("userId").(uuid.UUID)

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logrus.Errorf("error : %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user.Id = &userId

	updatedUser, err := users.UpdateUser(&user)
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
	body, _ := json.Marshal(updatedUser)
	_, _ = w.Write(body)
	return
}