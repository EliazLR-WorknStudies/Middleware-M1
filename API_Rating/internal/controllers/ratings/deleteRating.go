package ratings

import (
	"encoding/json"
	"middleware/ratings/internal/models"
	ratings "middleware/ratings/internal/services/ratings"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// DeleteRating
// @Tags         ratings
// @Summary      Delete a rating.
// @Description  Delete a rating.
// @Success      200            {object}  models.Ratings
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /ratings/{id} [delete]
func DeleteRating(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ratingId, _ := ctx.Value("id").(uuid.UUID)
	rating, err := ratings.DeleteRating(ratingId)

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
