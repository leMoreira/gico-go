package models

import "gorm.io/gorm"

type Pontos struct {
	gorm.Model
	Id       int `json:"id"`
	IdEquipe int `json:"idequipe"`
	IdProva  int `json:"idprova"`
	Pontos   int `json:"pontos"`
}
