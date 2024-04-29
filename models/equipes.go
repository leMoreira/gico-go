package models

import "gorm.io/gorm"

type Equipe struct {
	gorm.Model
	Id   int    `json:"id"`
	Nome string `json:"nome"`
	Cor  string `json:"cor"`
}
