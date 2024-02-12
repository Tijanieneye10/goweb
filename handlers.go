package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.gohtml")
}

func renderTemplate(w http.ResponseWriter, gohtml string) {
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatalln(err)
	}

	t, ok := tc[gohtml]

	if !ok {
		log.Fatalln(err)
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, nil)

	if err != nil {
		log.Println(err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

	parsedTemplate, _ := template.ParseFiles("./templates/"+gohtml, "./templates/base.layout.gohtml")

	err = parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing template", err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts

	}

	return myCache, nil

}
