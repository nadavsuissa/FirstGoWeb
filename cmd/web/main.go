package main

import (
	"FirstGoWeb/pkg/config"
	"FirstGoWeb/pkg/handlers"
	"FirstGoWeb/pkg/render"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("Error", err)
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Println(fmt.Sprintf("Starting on port %s", portNumber))
	_ = http.ListenAndServe(":8080", nil)

}
