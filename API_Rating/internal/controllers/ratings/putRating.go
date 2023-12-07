package ratings

import (
	"encoding/json"
	"middleware/ratings/internal/models"
	ratings "middleware/ratings/internal/services/ratings"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// PutRating
// @Tags         ratings
// @Summary      Update a rating.
// @Description  Update a rating.
// @Success      200            {object}  models.Collection
// @Failure      500             "Something went wrong"
// @Router       /ratings [get]
func PutRating(w http.ResponseWriter, r *http.Request) {
	// calling service
	ctx := r.Context()
	ratingId, _ := ctx.Value("id").(uuid.UUID)
	var rating models.Ratings
	rating.Id = &ratingId
	err := json.NewDecoder(r.Body).Decode(&rating)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ratings, err := ratings.UpdateRating(&rating)

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
