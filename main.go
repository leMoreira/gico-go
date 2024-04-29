package main

import (
	"github.com/leandro/gico-go/basedados"
	"github.com/leandro/gico-go/routes"
)

func main() {
	basedados.ConectaComBancoDeDados()

	routes.HandleRequests()

}
