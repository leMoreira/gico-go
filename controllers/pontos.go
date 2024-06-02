package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leandro/gico-go/basedados"
	"github.com/leandro/gico-go/models"
)

func ExibePontos(c *gin.Context) {
	var totalpontos models.TotalPontos

	type ResultadoPontos struct {
		EquipeNome  string
		TotalPontos int
	}

	type Results struct {
		Nome      string
		Cor       string
		Pontos    int
		Descricao string
		IdProva   int
		IdEquipe  int
	}

	list := []Results{}
	totalList := []ResultadoPontos{}

	query := basedados.DB.Table("pontos").
		Joins("INNER JOIN equipes ON equipes.id = pontos.id_equipe").
		Joins("INNER JOIN provas ON provas.id = pontos.id_prova").
		Select("pontos.id As idponto,pontos.deleted_at, pontos.pontos,pontos.id_prova,pontos.id_equipe, equipes.nome, equipes.cor, equipes.id As idequipe, provas.id As idprova, provas.descricao ").
		Find(&list)
	if query.Error != nil {
		panic(query.Error)
	}

	basedados.DB.Find(&totalpontos)

	basedados.DB.Table("equipes").
		Select("equipes.nome as equipe_nome, sum(pontos.pontos) as total_pontos").Joins("inner join pontos on equipes.id = pontos.id_equipe").Group("equipes.nome").Order("total_pontos desc").Scan(&totalList)

	fmt.Println(totalList)

	c.HTML(http.StatusOK, "exibepontos.html", gin.H{
		"list":      list,
		"listtotal": totalList,
	})

}

func DeletePontos(c *gin.Context) {

	var pontos []models.Pontos

	idprova := c.Query("idprova")
	idequipe := c.Query("idequipe")
	fmt.Println(idequipe)
	fmt.Println(idprova)
	basedados.DB.Where("id_equipe = ?", idequipe).
		Where("id_prova = ?", idprova).Unscoped().Delete(&pontos)

	c.HTML(http.StatusOK, "statusok.html", gin.H{
		"message": "Pontos Excluído",
		"link":    "/pontos",
		"status":  "text-success",
	})

}

func ExbitelaDeCadastro(c *gin.Context) {
	var equipes []models.Equipe
	var provas []models.Prova

	basedados.DB.Find(&equipes)
	basedados.DB.Find(&provas)

	c.HTML(http.StatusOK, "registerpontos.html", gin.H{
		"equipes": equipes,
		"provas":  provas,
	})

}
func RegistrarPontos(c *gin.Context) {

	var pontos models.Pontos

	DescricaoProva := c.PostForm("prova")
	NomeEquipe := c.PostForm("equipe")
	MarcarPontos := c.PostForm("pontos")

	pontos.IdProva, _ = strconv.Atoi(DescricaoProva)
	pontos.IdEquipe, _ = strconv.Atoi(NomeEquipe)
	pontos.Pontos, _ = strconv.Atoi(MarcarPontos)

	basedados.DB.Create(&pontos)

	c.HTML(http.StatusOK, "statusok.html", gin.H{
		"message": "Pontos Cadastrado",
		"link":    "/pontos",
		"status":  "text-success",
	})
}
