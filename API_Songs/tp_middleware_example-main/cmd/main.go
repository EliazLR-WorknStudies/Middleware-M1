package main

import (
	"middleware/example/internal/controllers/collections"
	"middleware/example/internal/helpers"
	_ "middleware/example/internal/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()

	r.Route("/songs", func(r chi.Router) {
		r.Get("/", collections.GetCollections)
		r.Post("/", collections.CreateCollection)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(collections.Ctx)
			r.Get("/", collections.GetCollection)
			r.Put("/", collections.UpdateCollection)
			r.Delete("/", collections.DeleteCollection)
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
		`CREATE TABLE IF NOT EXISTS songs (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			songName VARCHAR(255) NOT NULL,
			songAuthor VARCHAR(255) NOT NULL,
			songGenre VARCHAR(255) NOT NULL

		);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)
}
