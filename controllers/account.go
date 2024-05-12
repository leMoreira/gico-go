package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/leandro/gico-go/basedados"
	"github.com/leandro/gico-go/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", nil)

}

func Register(c *gin.Context) {

	c.HTML(http.StatusOK, "register.html", nil)

}

func FazerRegistroDeUsuario(c *gin.Context) {

	var usuario models.Usuario

	nome := c.PostForm("nome")
	email := c.PostForm("email")
	senha := c.PostForm("senha")

	senhaHash, err := bcrypt.GenerateFromPassword([]byte(senha), 14)
	if err != nil {
		fmt.Println("Ops Algo deu errado com o Hash")
	}

	usuario.Nome = nome
	usuario.Email = email
	usuario.Senha = string(senhaHash)

	basedados.DB.Create(&usuario)
	c.HTML(http.StatusOK, "statusok.html", gin.H{
		"message": "Usuário cadastrado ... ",
		"link":    "login",
	})
}

func Logarusuario(c *gin.Context) {
	var usuario models.Usuario

	email := c.PostForm("email")
	senha := c.PostForm("senha")

	basedados.DB.First(&usuario, "email = ?", email)
	if usuario.Id == 0 {
		c.HTML(http.StatusBadRequest, "statusok.html", gin.H{
			"message": "Usuário não encontrado ... ",
			"link":    "login",
			"status":  "text-danger",
		})

		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(usuario.Senha), []byte(senha))

	if err != nil {
		c.HTML(http.StatusBadRequest, "statusok.html", gin.H{
			"message": "Usuário não encontrado ... ",
			"link":    "login",
			"status":  "text-danger",
		})

		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": usuario.Id,
		"exp": time.Now().Add(time.Hour * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("Lm0712"))

	if err != nil {
		c.HTML(http.StatusBadRequest, "statusok.html", gin.H{
			"message": "Não criou o token ... ",
			"link":    "/login",
			"status":  "text-danger",
		})

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*30, "", "", false, true)

	c.HTML(http.StatusOK, "statusok.html", gin.H{
		"message": "Entrando .. ",
		"link":    "/",
		"token":   tokenString,
		"status":  "text-success",
	})
}

func Validate(c *gin.Context) {

	usuario, _ := c.Get("usuario")

	c.HTML(http.StatusOK, "index.html", gin.H{
		"message": usuario,
	})

}

func DeleteCookie(c *gin.Context) {

	c.SetCookie("Authorization", "usuario", -1, "/", "", true, false)
	c.HTML(http.StatusOK, "statusok.html", gin.H{
		"message": "Saindo ... ",
		"link":    "/login",
		"status":  "text-dark",
	})

}
