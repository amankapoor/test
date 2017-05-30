package common

import (
	"html/template"
	"log"
	"net/http"
)

//RenderTemplate is a Template rendering function to avoid duplication of code in each handler
func RenderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {

	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal("Error encountered while parsing the template: ", err)
	}
	t.Execute(w, templateData)
}
