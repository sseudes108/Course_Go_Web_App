package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sseudes108/go-course/pkg/config"
	"github.com/sseudes108/go-course/pkg/handlers"
	"github.com/sseudes108/go-course/pkg/render"
)

const portNumber = ":8080"

// main is the Main aplication function
func main() {
	var app config.AppConfig

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting aplication on port%s", portNumber))

	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = server.ListenAndServe()
	log.Fatal(err)
}
