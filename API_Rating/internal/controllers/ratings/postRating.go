package ratings

import (
	"encoding/json"
	"middleware/ratings/internal/models"
	ratings "middleware/ratings/internal/services/ratings"
	"net/http"

	"github.com/sirupsen/logrus"
)

// PostRating
// @Tags         ratings
// @Summary      Create a rating.
// @Description  Create a rating.
// @Success      200            {object}  models.Collection
// @Failure      500             "Something went wrong"
// @Router       /ratings [get]
func PostRating(w http.ResponseWriter, r *http.Request) {
	// calling service
	var rating models.Ratings

	err := json.NewDecoder(r.Body).Decode(&rating)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ratings, err := ratings.CreateRating(&rating)

	//ERRORS TO DO ?
	if err != nil {
		// logging error
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			// writing http code in header
			w.WriteHeader(customError.Code)
			// writing error message in body
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(ratings)
	_, _ = w.Write(body)
	return
}
