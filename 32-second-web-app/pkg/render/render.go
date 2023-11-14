package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	parsedTemplate, err := template.ParseFiles("templates/layout.html", "templates/"+templateName+".html")
	if err != nil {
		panic(err)
	}
	err = parsedTemplate.Execute(w, data)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}
