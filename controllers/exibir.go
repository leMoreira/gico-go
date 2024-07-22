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

func DetalhesPontuacaousuario(c *gin.Context) {
	type PontosEADescricao struct {
		EquipeNome     string `db:"equipe_nome"`
		DescricaoProva string `db:"descricao_prova"`
		TotalPontos    int    `db:"total_pontos"`
	}

	pontosDescricao := []PontosEADescricao{}

	equipenome := c.Request.URL.Query().Get("nomeequipe")

	basedados.DB.Table("pontos").
		Select("pontos.pontos as total_pontos, equipes.nome as equipe_nome, provas.descricao as descricao_prova").
		Joins("INNER JOIN provas ON pontos.id_prova= provas.id").
		Joins("INNER JOIN equipes ON pontos.id_equipe = equipes.id").
		Where("equipes.nome = ?", equipenome).
		Find(&pontosDescricao)

	// implementando

	c.HTML(http.StatusOK, "detalhesprova.html", gin.H{
		"provasdescricao": pontosDescricao,
		"nomedaequipe":    equipenome,
	})

}
