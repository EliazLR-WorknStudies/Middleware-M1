package collections

import (
	"encoding/json"
	"middleware/example/internal/models"
	"middleware/example/internal/services/collections"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// UpdateCollection
// @Tags         songs
// @Summary      Update a collection.
// @Description  Update a collection.
// @Param        id            	path      string  	   true  "Collection UUID formatted ID"
// @Param        body          {object}   models.Song  true  "Song object"
// @Success      200            {object}  models.Song
// @Failure      500            "Something went wrong"
// @Router       /collections/{id} [put]
func UpdateCollection(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	collectionId, _ := ctx.Value("collectionId").(uuid.UUID)

	var songData models.Song
	err := json.NewDecoder(r.Body).Decode(&songData)

	if err != nil {
		// Gérez les erreurs de décodage JSON
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = collections.PutCollectionById(collectionId, songData.SongName, songData.SongAuthor, songData.SongGenre)
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
	w.Write([]byte("Song updated successfully"))
}
