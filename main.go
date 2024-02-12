package main

import (
	"github.come/Tijanieneye10/goweb/pkg/config"
	"net/http"
)

func main() {
	var app config.AppConfig
	tc, err := createTemplateCache()
	if err != nil {
		return
	}

	app.TemplateCache = tc

	NewTemplates(&app)

	http.HandleFunc("/", home)

	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		return
	}

}
