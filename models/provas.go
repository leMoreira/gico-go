package models

import "gorm.io/gorm"

type Prova struct {
	gorm.Model
	Id              int    `json:"id"`
	Descrição       string `json:"descricao"`
	pontos_primeiro int    `json:"pontos_primeiro"`
	pontos_segundo  int    `json:"pontos_segundo"`
	pontos_terceiro int    `json:"pontos_terceiro"`
	pontos_quarto   int    `json:prontos_quarto`
}
