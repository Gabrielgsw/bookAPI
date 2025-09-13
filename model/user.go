package model

type User struct {
	Id    string `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}
