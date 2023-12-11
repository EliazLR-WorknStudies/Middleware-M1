package models

import (
	"github.com/gofrs/uuid"
)

type Song struct {
	Id         *uuid.UUID `json:"id"`
	SongName   string     `json:"songname"`
	SongAuthor string     `json:"songauthor"`
	SongGenre  string     `json:"songgenre"`
}
