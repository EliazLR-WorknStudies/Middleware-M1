package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"middleware/users/internal/controllers/users"
	"middleware/users/internal/helpers"
	_ "middleware/users/internal/models"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) {
		r.Get("/", users.GetUsers)
		r.Post("/", users.CreateUser)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(users.Ctx)
			r.Get("/", users.GetUser)
			r.Delete("/", users.DeleteUser)
			r.Put("/", users.UpdateUser)
		})
	})

	logrus.Info("[INFO] Web server started. Now listening on *:8081")
	logrus.Fatalln(http.ListenAndServe(":8081", r))
}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	schemes := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			username VARCHAR(255) NOT NULL
		);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)
}
