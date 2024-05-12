package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Id    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}
