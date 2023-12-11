package collections

import (
	"context"
	"encoding/json"
	"fmt"
	"middleware/example/internal/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// Ici la fonction Ctx fait la requête et nous permet de savoir si on peut récupérer l'ID

func Ctx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		collectionId, err := uuid.FromString(chi.URLParam(r, "id"))
		if err != nil {
			logrus.Errorf("parsing error : %s", err.Error())
			customError := &models.CustomError{
				Message: fmt.Sprintf("cannot parse id (%s) as UUID", chi.URLParam(r, "id")),
				Code:    http.StatusUnprocessableEntity,
			}
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
			return

		}

		ctx := context.WithValue(r.Context(), "collectionId", collectionId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}