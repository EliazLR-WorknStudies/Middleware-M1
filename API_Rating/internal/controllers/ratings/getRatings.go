package ratings

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"middleware/ratings/internal/models"
	ratings "middleware/ratings/internal/services/ratings"
	"net/http"
)

// Getratings
// @Tags         ratings
// @Summary      Get ratings.
// @Description  Get ratings.
// @Success      200            {array}  models.Collection
// @Failure      500             "Something went wrong"
// @Router       /ratings [get]
func GetRatings(w http.ResponseWriter, _ *http.Request) {
	// calling service
	ratings, err := ratings.GetAllRatings()

	//ERRORS
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
