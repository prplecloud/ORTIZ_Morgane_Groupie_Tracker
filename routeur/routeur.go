package routeur

import (
	"fmt"
	"net/http"
	"pokeapi/controller"
)

func Initserv() {
	css := http.FileServer(http.Dir("./assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", css))

	http.HandleFunc("/home", controller.Home)
	http.HandleFunc("/collection", controller.RenderPokemonPage)
	http.HandleFunc("/pokemon/", controller.PokeDetails)
	http.HandleFunc("/categorie", controller.Pokemon) //=Pokemon
	http.HandleFunc("/favoris", controller.ViewFavoritesHandler)
	http.HandleFunc("/add-to-favorites", controller.AddToFavoritesHandler)
	http.HandleFunc("/about", controller.Home)
	http.HandleFunc("/result", controller.Search)
	http.HandleFunc("/filter", controller.FilterHandler)
	http.HandleFunc("/error", controller.Search)

	fmt.Println("serveur ouvert sur le port 8080")
	http.ListenAndServe(":8080", nil)
}
