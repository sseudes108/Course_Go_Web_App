package main

import (
	"fmt"
	"net/http"

	"github.com/sseudes108/go-course/pkg/handlers"
)

const portNumber = ":8108"

// main is the Main aplication function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting aplication on port%s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
