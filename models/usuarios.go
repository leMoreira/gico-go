package models

type Usuario struct {
	Id    int    `json:"id"`
	Nome  string `json:"nome"`
	email string `json:"email"`
	senha string `json:"senha"`
}
