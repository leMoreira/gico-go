package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", nil)

}

func Register(c *gin.Context) {

	c.HTML(http.StatusOK, "register.html", nil)

}
