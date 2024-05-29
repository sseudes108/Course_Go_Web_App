package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sseudes108/go-course/pkg/config"
	"github.com/sseudes108/go-course/pkg/handlers"
	"github.com/sseudes108/go-course/pkg/render"
)

const portNumber = ":8108"

// main is the Main aplication function
func main() {
	var app config.AppConfig

	templateCache, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = templateCache

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting aplication on port%s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
