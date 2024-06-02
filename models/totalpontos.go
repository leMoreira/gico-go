package models

import "gorm.io/gorm"

type TotalPontos struct {
	gorm.Model
	Idequipe    int `jason:"idequipe"`
	TotalPontos int `jason:"totalpontos"`
}
