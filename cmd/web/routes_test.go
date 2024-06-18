package main

import (
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/sseudes108/Course_Go_Web_App/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch mux.(type) {
	case *chi.Mux:
		//Pass - Type is correct
	default:
		t.Error("Type is not chi.Mux")
	}
}
