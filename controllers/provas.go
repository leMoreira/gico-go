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

	prova.Descricao = descricao

	basedados.DB.Create(&prova)

	c.HTML(http.StatusOK, "statusok.html", gin.H{
		"message": "Prova Cadastrada",
		"link":    "/provas",
	})

}

func DeletaProva(c *gin.Context) {
	var prova []models.Prova

	id := c.Query("id")

	basedados.DB.Delete(&prova, id)

	c.HTML(http.StatusOK, "statusok.html", gin.H{
		"message": "Prova Excluída",
		"link":    "/provas",
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

	id := c.PostForm("id")

	idConvert, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Não foi convertido")
		log.Panic(err.Error())
	}

	basedados.DB.Model(&Prova).Where("id = ?", idConvert).Updates(models.Prova{
		Descricao: descricao,
	})

	fmt.Println(&Prova)

	c.HTML(http.StatusOK, "statusok.html", gin.H{
		"message": "Prova Atualizada",
		"link":    "/provas",
		"status":  "text-success",
	})

}
