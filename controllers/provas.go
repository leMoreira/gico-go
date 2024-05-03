package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leandro/gico-go/basedados"
	"github.com/leandro/gico-go/models"
)

func ExibeProvas(c *gin.Context) {

	c.HTML(http.StatusOK, "exibeprovas.html", nil)

}

func InserirDadosProva(c *gin.Context) {

	var prova models.Prova

	descricao := c.PostForm("descricao")
	pontosprimeiro := c.PostForm("pontosprimeiro")
	pontossegundo := c.PostForm("pontossegundo")
	pontosterceiro := c.PostForm("pontosterceiro")
	pontosquarto := c.PostForm("pontosquarto")

	prova.Descrição = descricao
	prova.Pontosprimeiro, _ = strconv.Atoi(pontosprimeiro)
	prova.Pontossegundo, _ = strconv.Atoi(pontossegundo)
	prova.Pontosterceiro, _ = strconv.Atoi(pontosterceiro)
	prova.Pontosquarto, _ = strconv.Atoi(pontosquarto)

	basedados.DB.Create(&prova)

	c.HTML(http.StatusOK, "statusok.html", gin.H{
		"message": "Prova Cadastrada",
	})

}
