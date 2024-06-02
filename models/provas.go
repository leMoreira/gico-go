package models

import "gorm.io/gorm"

type Prova struct {
	gorm.Model
	Id        int    `json:"id"`
	Descricao string `json:"descricao"`
}
