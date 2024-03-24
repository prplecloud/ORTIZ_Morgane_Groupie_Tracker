package templates

import (
	"fmt"
	"html/template"
)

var Temp *template.Template

func InitTemplate() {

	temp, errTemp := template.ParseGlob("./templates/*.html")
	if errTemp != nil {
		fmt.Printf("Erreur template: %v", errTemp.Error())
		return
	}
	Temp = temp

}
