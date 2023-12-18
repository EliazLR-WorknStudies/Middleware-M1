package users

import (
	"middleware/users/internal/helpers"
	"middleware/users/internal/models"

	"github.com/gofrs/uuid"
)

func GetAllUsers() ([]models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM users")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	users := []models.User{}
	for rows.Next() {
		var data models.User
		err = rows.Scan(&data.Id, &data.Username)
		if err != nil {
			return nil, err
		}
		users = append(users, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return users, err
}

func GetUserById(id uuid.UUID) (*models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM users WHERE id=?", id.String())
	helpers.CloseDB(db)

	var user models.User
	err = row.Scan(&user.Id, &user.Username)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func CreateUser(user *models.User) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO users (id, username) VALUES (?, ?)", user.Id, user.Username)
	helpers.CloseDB(db)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM users WHERE id=?", id)
	helpers.CloseDB(db)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *models.User) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE users SET username=? WHERE id=?", user.Username, user.Id)
	helpers.CloseDB(db)
	if err != nil {
		return err
	}
	return nil
}
