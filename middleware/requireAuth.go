package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/leandro/gico-go/basedados"
	"github.com/leandro/gico-go/models"
)

var secret = []byte("Lm0712")

func RequireAuth(c *gin.Context) {
	// Get the cookie off the request
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.HTML(http.StatusUnauthorized, "login.html", nil)
	}

	// Decode/validate it
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Chec k the expiry date
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			c.HTML(http.StatusUnauthorized, "login.html", nil)
		}

		// Find the user with token Subject
		var usuario models.Usuario
		basedados.DB.First(&usuario, claims["sub"])

		if usuario.Id == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			c.HTML(http.StatusUnauthorized, "login.html", nil)
		}

		// Attach the request
		c.Set("usuario", usuario)

		//Continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.HTML(http.StatusUnauthorized, "login.html", nil)
	}

}
