package models

type UserJSON struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type TokenJSON struct {
	Token string `json:"token"`
}
