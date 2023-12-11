package collections

import (
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"

	"github.com/gofrs/uuid"
)

func GetAllCollections() ([]models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM songs")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	collections := []models.Song{}
	for rows.Next() {
		var data models.Song
		err = rows.Scan(&data.Id, &data.SongName, &data.SongAuthor, &data.SongGenre)
		if err != nil {
			return nil, err
		}
		collections = append(collections, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return collections, err
}

func GetCollectionById(id uuid.UUID) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM songs WHERE id=?", id.String())
	helpers.CloseDB(db)

	var collection models.Song
	err = row.Scan(&collection.Id, &collection.SongName, &collection.SongAuthor, &collection.SongGenre)
	if err != nil {
		return nil, err
	}
	return &collection, err
}

func PutCollectionById(id uuid.UUID, songName string, songAuthor string, songGenre string) (*models.Song, error) {

	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("UPDATE songs SET songName = ?, songAuthor = ?, songGenre = ? WHERE id = ?", songName, songAuthor, songGenre, id.String())
	helpers.CloseDB(db)
	//fmt.Print(resultat)
	var collection models.Song
	//row.Scan(&collection.Id, &collection.SongName)
	if err != nil {
		return nil, err
	}
	return &collection, err
}

func CreateCollectionByRepo(id uuid.UUID, songName string, songAuthor string, songGenre string) (*models.Song, error) {

	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("INSERT INTO songs (id, songName, songAuthor, songGenre) VALUES (?, ?, ?, ?)", id.String(), songName, songAuthor, songGenre)
	helpers.CloseDB(db)

	var collection models.Song
	//row.Scan(&collection.Id, &collection.SongName)
	if err != nil {
		return nil, err
	}
	return &collection, err
}

func DeleteCollectionByRepo(id uuid.UUID) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("DELETE FROM songs WHERE id = ?", id.String())
	helpers.CloseDB(db)

	var collection models.Song
	if err != nil {
		return nil, err
	}
	return &collection, err

}

/*
func UpdateCollectionById(id uuid.UUID, songname string) (*models.Collection, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	row := db.QueryRow("UPDATE collections SET songName = ? WHERE id = ?", songname, id.String())
	helpers.CloseDB(db)

	var collection models.Collection
	err = row.Scan(&collection.Id, &collection.SongName)
	if err != nil {
		return nil, err
	}
	return &collection, err
}
*/
