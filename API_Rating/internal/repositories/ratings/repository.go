package ratings

import (
	"github.com/gofrs/uuid"
	"middleware/ratings/internal/helpers"
	"middleware/ratings/internal/models"
	"strconv"
	"fmt"
)

func GetAllRatings() ([]models.Ratings, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM ratings")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	ratings := []models.Ratings{}
	for rows.Next() {
		var data models.Ratings
		err = rows.Scan(&data.Id, &data.IdSong, &data.IdUser, &data.Comment, &data.Rating)
		if err != nil {
			return nil, err
		}
		ratings = append(ratings, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return ratings, err
}


func GetRatingById(id uuid.UUID) (*models.Ratings, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	fmt.Printf(id.String())
	row := db.QueryRow("SELECT * FROM ratings WHERE id=?", id.String())
	helpers.CloseDB(db)

	var rating models.Ratings
	err = row.Scan(&rating.Id, &rating.IdSong, &rating.IdUser, &rating.Comment, &rating.Rating)
	if err != nil {
		return nil, err
	}
	return &rating, err
}

func CreateRating(rating *models.Ratings) (*models.Ratings, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	
	db.Exec("INSERT INTO ratings(id,idSong,idUser,comment,rating) VALUES(?,?,?,?,?);",rating.Id.String(),rating.IdSong.String(),rating.IdUser.String(),rating.Comment,strconv.Itoa(rating.Rating))
	
	helpers.CloseDB(db)

	return rating, err
}