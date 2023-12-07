package collections

import (
	"encoding/json"
	"fmt"
	"middleware/example/internal/models"
	"middleware/example/internal/services/collections"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// CreateCollection
// @Tags         collections
// @Summary      Create a collection.
// @Description  Create a collection.
// @Param        id            	path      string  	   true  "Collection UUID formatted ID"
// @Param        body          {object}   models.Song  true  "string"
// @Success      200            {object}  models.Song
// @Failure      500            "Something went wrong"
// @Router       /collections/{id} [post]
func CreateCollection(w http.ResponseWriter, r *http.Request) {

	var songData models.Song
	err := json.NewDecoder(r.Body).Decode(&songData)
	if err != nil {
		// Gérez les erreurs de décodage JSON
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newID, err := uuid.NewV4()
	if err != nil {
		// Gérer l'erreur lors de la génération de l'ID
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Print(songData)
	_, err = collections.PostCollectionById(newID, songData.SongName, songData.SongAuthor, songData.SongName)
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
	w.Write([]byte("Song created successfully"))

}
