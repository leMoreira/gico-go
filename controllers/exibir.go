package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leandro/gico-go/basedados"
)

func ExibirPontosusuarios(c *gin.Context) {

	type ResultadoPontos struct {
		EquipeNome  string `db:"equipe_nome"`
		EquipeCor   string `db:"equipe_cor"`
		TotalPontos int    `db:"total_pontos"`
	}

	totalList := []ResultadoPontos{}

	basedados.DB.Table("equipes").
		Select(" equipes.nome as equipe_nome, equipes.cor as equipe_cor, sum(pontos.pontos) as total_pontos").Joins("inner join pontos on equipes.id = pontos.id_equipe").Group("equipes.nome, equipes.cor").Order("total_pontos desc").Scan(&totalList)

	c.HTML(http.StatusOK, "users.html", gin.H{
		"totalLista": totalList,
	})
}

func DetalhesPontuacaousuario(w http.ResponseWriter, r *http.Request) {

	type EquipePontos struct {
		EquipeNome     string `db:"equipe_nome"`
		ProvaEquipe    string `db:"prova_equipe"`
		PontosPorProva int    `db:"pontos_prova"`
	}

	totalPontosPorEquipe := []EquipePontos{}

	basedados.DB.Table("equipes").
		Select(" equipes.nome as equipe_nome, sum(pontos.pontos) as total_pontos").
		Joins("inner join pontos on equipes.id = pontos.id_equipe").
		Group("equipes.nome, equipes.cor").Order("total_pontos desc").
		Scan(&totalPontosPorEquipe)

}
