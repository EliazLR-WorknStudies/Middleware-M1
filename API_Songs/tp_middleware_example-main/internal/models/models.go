package models

import (
	"github.com/gofrs/uuid"
)

type Collection struct {
	Id       *uuid.UUID `json:"id"`
	SongName string     `json:"songname"`
}
