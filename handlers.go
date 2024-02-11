package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.gohtml")
}

func renderTemplate(w http.ResponseWriter, gohtml string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+gohtml, "./templates/base.layout.gohtml")

	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing template", err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
}
