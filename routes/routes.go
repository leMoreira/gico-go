package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/leandro/gico-go/controllers"
)

func HandleRequests() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.GET("/", controllers.Saudacao)
	r.GET("/equipes", controllers.ExibeEquipes)
	r.POST("/insert", controllers.InserirDadosEquipes)
	r.GET("/tudocerto", controllers.StatusDeInsercaoCorreta)
	r.GET("/delete", controllers.DeletaEquipe)
	r.GET("/edit", controllers.BuscarEquipeParaEditar)
	r.POST("/update", controllers.AtualizarEquipe)

	r.Run()
}
