package models

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type UserID int

type UserJSON struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type TokenJSON struct {
	Token string `json:"token"`
}

type NoteJSON struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type SpellerJSON struct {
	Code int      `json:"code"`
	Pos  int      `json:"pos"`
	Row  int      `json:"row"`
	Col  int      `json:"col"`
	Len  int      `json:"len"`
	Word string   `json:"word"`
	S    []string `json:"s"`
}
