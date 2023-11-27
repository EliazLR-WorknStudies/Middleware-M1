package ratings

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/ratings/internal/models"
	ratings "middleware/ratings/internal/services/ratings"
	"net/http"
)

// GetRating
// @Tags         ratings
// @Summary      Get a rating.
// @Description  Get a rating.
// @Param        id           	path      string  true  "Collection UUID formatted ID"
// @Success      200            {object}  models.Collection
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /ratings/{id} [get]
func GetRating(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ratingId, _ := ctx.Value("id").(uuid.UUID)

	rating, err := ratings.GetRatingById(ratingId)
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
	body, _ := json.Marshal(rating)
	_, _ = w.Write(body)
	return
}
