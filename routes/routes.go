package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/leandro/gico-go/controllers"
	"github.com/leandro/gico-go/middleware"
)

func HandleRequests() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	//Login
	r.GET("/login", controllers.Login)
	r.GET("/register", controllers.Register)
	r.POST("/registrarusuario", controllers.FazerRegistroDeUsuario)
	r.POST("/logar", controllers.Logarusuario)
	r.GET("/deleteCookie", controllers.DeleteCookie)

	//Equipes
	r.GET("/admin", middleware.RequireAuth, controllers.Index)
	r.GET("/equipes", middleware.RequireAuth, controllers.ExibeEquipes)
	r.POST("/equipes/insert", middleware.RequireAuth, controllers.InserirDadosEquipes)
	r.GET("/tudocerto", middleware.RequireAuth, controllers.StatusDeInsercaoCorreta)
	r.GET("/equipes/delete", middleware.RequireAuth, controllers.DeletaEquipe)
	r.GET("/equipes/edit", middleware.RequireAuth, controllers.BuscarEquipeParaEditar)
	r.POST("/equipes/update", middleware.RequireAuth, controllers.AtualizarEquipe)

	// Provas
	r.GET("/provas", middleware.RequireAuth, controllers.ExibeProvas)
	r.POST("provas/insert", middleware.RequireAuth, controllers.InserirDadosProva)
	r.GET("/provas/edit", middleware.RequireAuth, controllers.BuscarProvaParaEditar)
	r.GET("/provas/delete", middleware.RequireAuth, controllers.DeletaProva)
	r.POST("/provas/update", middleware.RequireAuth, controllers.AtualizarProva)

	// Pontos

	r.GET("/pontos", middleware.RequireAuth, controllers.ExibePontos)
	r.GET("/pontos/delete", middleware.RequireAuth, controllers.DeletePontos)
	r.GET("/pontos/registerpontos", middleware.RequireAuth, controllers.ExbitelaDeCadastro)
	r.POST("/pontos/cadastrarpontos", middleware.RequireAuth, controllers.RegistrarPontos)

	//front

	r.GET("/", controllers.ExibirPontosusuarios)

	r.Run("0.0.0.0:80")
}
