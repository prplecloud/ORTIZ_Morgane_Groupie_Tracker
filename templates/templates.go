package templates

import (
	"fmt"
	"html/template"
)

func InitTemplate() {
	_, err := template.ParseGlob("./templates/*.html")
	if err != nil {
		fmt.Printf(fmt.Sprintf("ERREUR => %s", err.Error()))
		return
	}
}
