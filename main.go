package main

import (
	"pokeapi/routeur"
	"pokeapi/templates"
)

func main() {

	templates.InitTemplate()
	routeur.Initserv()
}
