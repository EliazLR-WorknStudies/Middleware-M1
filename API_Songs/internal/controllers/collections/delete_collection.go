package collections

import (
	"encoding/json"
	"middleware/example/internal/models"
	"middleware/example/internal/services/collections"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// DeleteCollection
// @Tags         songs
// @Summary      Delete a song.
// @Description  Delete a song.
// @Param        id           	path      string  true  "Collection UUID formatted ID"
// @Success      204            {object}  models.Song
// @Failure      500            "Something went wrong"
// @Router       /collections [delete]
func DeleteCollection(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	collectionId, _ := ctx.Value("collectionId").(uuid.UUID)

	_, err := collections.DeleteCollectionById(collectionId)
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
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Song deleted successfully"))
}
