package models

import (
	"github.com/gofrs/uuid"
)

type Collection struct {
	Id      *uuid.UUID `json:"id"`
	Content string     `json:"content"`
}

type User struct {
	Id       *uuid.UUID `json:"id"`
	Username string     `json:"username"`
}
