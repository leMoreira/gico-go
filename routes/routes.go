package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/leandro/gico-go/controllers"
)

func HandleRequests() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	//Equipes
	r.GET("/", controllers.Index)
	r.GET("/equipes", controllers.ExibeEquipes)
	r.POST("/equipes/insert", controllers.InserirDadosEquipes)
	r.GET("/tudocerto", controllers.StatusDeInsercaoCorreta)
	r.GET("/equipes/delete", controllers.DeletaEquipe)
	r.GET("/equipes/edit", controllers.BuscarEquipeParaEditar)
	r.POST("/equipes/update", controllers.AtualizarEquipe)

	// Provas
	r.GET("/provas", controllers.ExibeProvas)
	r.POST("provas/insert", controllers.InserirDadosProva)

	r.Run()
}
