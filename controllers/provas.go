package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leandro/gico-go/basedados"
	"github.com/leandro/gico-go/models"
)

func ExibeProvas(c *gin.Context) {

	var provas []models.Prova
	basedados.DB.Find(&provas)

	c.HTML(http.StatusOK, "exibeprovas.html", gin.H{
		"provas": provas,
	})

}

func InserirDadosProva(c *gin.Context) {

	var prova models.Prova

	descricao := c.PostForm("descricao")
	pontosprimeiro := c.PostForm("pontosprimeiro")
	pontossegundo := c.PostForm("pontossegundo")
	pontosterceiro := c.PostForm("pontosterceiro")
	pontosquarto := c.PostForm("pontosquarto")

	prova.Descricao = descricao
	prova.Pontosprimeiro, _ = strconv.Atoi(pontosprimeiro)
	prova.Pontossegundo, _ = strconv.Atoi(pontossegundo)
	prova.Pontosterceiro, _ = strconv.Atoi(pontosterceiro)
	prova.Pontosquarto, _ = strconv.Atoi(pontosquarto)

	basedados.DB.Create(&prova)

	c.HTML(http.StatusOK, "statusok.html", gin.H{
		"message": "Prova Cadastrada",
		"link":    "provas",
	})

}

func DeletaProva(c *gin.Context) {
	var prova []models.Prova

	id := c.Query("id")

	basedados.DB.Delete(&prova, id)

	c.HTML(http.StatusOK, "statusok.html", gin.H{
		"message": "Prova Excluída",
		"link":    "provas",
	})

}

func BuscarProvaParaEditar(c *gin.Context) {
	var prova []models.Prova

	id := c.Query("id")

	basedados.DB.First(&prova, id)

	c.HTML(http.StatusOK, "editaprova.html", gin.H{
		"prova": prova,
	})

}

func AtualizarProva(c *gin.Context) {
	var Prova models.Prova
	descricao := c.PostForm("descricao")
	pontosprimeiro := c.PostForm("pontosprimeiro")
	pontossegundo := c.PostForm("pontossegundo")
	pontosterceiro := c.PostForm("pontosterceiro")
	pontosquarto := c.PostForm("pontosquarto")
	id := c.PostForm("id")

	idConvert, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Não foi convertido")
		log.Panic(err.Error())
	}

	pontosPrimeiroConvert, err := strconv.Atoi(pontosprimeiro)
	if err != nil {
		fmt.Println("Não foi convertido")
		log.Panic(err.Error())
	}
	pontosSegundoConvert, err := strconv.Atoi(pontossegundo)
	if err != nil {
		fmt.Println("Não foi convertido")
		log.Panic(err.Error())
	}
	pontosTerceiroConvert, err := strconv.Atoi(pontosterceiro)
	if err != nil {
		fmt.Println("Não foi convertido")
		log.Panic(err.Error())
	}
	pontosQuartoConvert, err := strconv.Atoi(pontosquarto)
	if err != nil {
		fmt.Println("Não foi convertido")
		log.Panic(err.Error())
	}

	basedados.DB.Model(&Prova).Where("id = ?", idConvert).Updates(models.Prova{
		Descricao:      descricao,
		Pontosprimeiro: pontosPrimeiroConvert,
		Pontossegundo:  pontosSegundoConvert,
		Pontosterceiro: pontosTerceiroConvert,
		Pontosquarto:   pontosQuartoConvert,
	})

	fmt.Println(&Prova)

	c.HTML(http.StatusOK, "statusok.html", gin.H{
		"message": "Prova Atualizada",
		"link":    "provas",
	})

}
