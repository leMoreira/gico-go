package models

import "gorm.io/gorm"

type Prova struct {
	gorm.Model
	Id             int    `json:"id"`
	Descricao      string `json:"descricao"`
	Pontosprimeiro int    `json:"pontosprimeiro"`
	Pontossegundo  int    `json:"pontossegundo"`
	Pontosterceiro int    `json:"pontosterceiro"`
	Pontosquarto   int    `json:"pontosquarto"`
}
