package routeur

import (
	"fmt"
	"net/http"
)

func Initserv() {
	css := http.FileServer(http.Dir("./assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", css))

	
	fmt.Println("serveur ouvert sur le port 8080")
	http.ListenAndServe(":8080", nil)
}
