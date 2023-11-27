package models

import (
	"github.com/gofrs/uuid"
)

type Ratings struct {
	Id      *uuid.UUID `json:"id"`
	IdSong      *uuid.UUID `json:"idSong"`
	IdUser      *uuid.UUID `json:"idUser"`
	Comment 	string     `json:"comment"`
	Rating 		int     `json:"rating"`
}
