package main

import (
	"middleware/ratings/internal/controllers/ratings"
	"middleware/ratings/internal/helpers"
	_ "middleware/ratings/internal/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()

	r.Route("/ratings", func(r chi.Router) {
		r.Get("/", ratings.GetRatings)
		r.Post("/", ratings.PostRating)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(ratings.Ctx)
			r.Get("/", ratings.GetRating)
			r.Put("/", ratings.PutRating)
			r.Delete("/", ratings.DeleteRating)
		})
	})

	r.Route("/song", func(r chi.Router) {
		r.Route("/{id}", func(r chi.Router) {
			r.Use(ratings.Ctx)
			r.Get("/", ratings.GetRatingsFromSong)
		})
	})

	logrus.Info("[INFO] Web server started. Now listening on *:8080")
	logrus.Fatalln(http.ListenAndServe(":8080", r))
}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	schemes := []string{
		`CREATE TABLE IF NOT EXISTS ratings (
			id VARCHAR(255) PRIMARY KEY NOT NULL,
			idSong VARCHAR(255) NOT NULL,
			idUser VARCHAR(255) NOT NULL,
			comment VARCHAR(255) NOT NULL,
			rating VARCHAR(255) NOT NULL
		);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)
}
