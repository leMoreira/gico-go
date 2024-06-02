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

func Index(c *gin.Context) {

	usuario, _ := c.Get("usuario")

	fmt.Println(usuario)

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"usuario": usuario,
	})

}

func ExibeEquipes(c *gin.Context) {

	var equipes []models.Equipe
	basedados.DB.Find(&equipes)
	// Select * From equipes

	c.HTML(http.StatusOK, "exibeequipes.html", gin.H{
		"equipes": equipes,
	})
}

func InserirDadosEquipes(c *gin.Context) {

	var equipe models.Equipe
	nome := c.PostForm("nome")
	cor := c.PostForm("cor")

	equipe.Nome = nome
	equipe.Cor = cor

	basedados.DB.Create(&equipe)
	c.HTML(http.StatusOK, "statusok.html", gin.H{
		"message": "Equipe Cadastrada",
		"link":    "/equipes",
		"status":  "text-success",
	})
}

func StatusDeInsercaoCorreta(c *gin.Context) {

	c.HTML(http.StatusOK, "statusok.html", nil)

}

func DeletaEquipe(c *gin.Context) {

	var equipe []models.Equipe
	id := c.Query("id")

	basedados.DB.Delete(&equipe, id)

	c.HTML(http.StatusOK, "statusok.html", gin.H{
		"message": "Equipe Excluída",
		"link":    "/equipes",
		"status":  "text-success",
	})

}

func BuscarEquipeParaEditar(c *gin.Context) {

	var equipe []models.Equipe

	id := c.Query("id")

	basedados.DB.First(&equipe, id)

	c.HTML(http.StatusOK, "editarequipe.html", gin.H{
		"equipe": equipe,
	})

}

func AtualizarEquipe(c *gin.Context) {

	var Equipe models.Equipe
	nome := c.PostForm("nome")
	cor := c.PostForm("cor")
	id := c.PostForm("id")

	// basedados.DB.First(&Equipe, id)
	fmt.Println(nome)
	fmt.Println(cor)
	fmt.Println(id)
	idConvert, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Não foi convertido, ")
		log.Panic(err.Error())
	}
	// Equipe.Cor = cor
	// Equipe.Nome = nome

	basedados.DB.Model(&Equipe).Where("id = ?", idConvert).Updates(models.Equipe{
		Nome: nome,
		Cor:  cor,
	})

	fmt.Println(&Equipe)

	c.HTML(http.StatusOK, "statusok.html", gin.H{
		"message": "Equipe Atualizada",
		"link":    "/equipes",
		"status":  "text-success",
	})

}
