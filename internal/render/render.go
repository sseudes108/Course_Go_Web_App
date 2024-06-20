package render

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/sseudes108/Course_Go_Web_App/internal/config"
	"github.com/sseudes108/Course_Go_Web_App/internal/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

var pathToTemplates = "./templates"

// NewTemplate sets the config for the template package
func NewTemplates(appConfig *config.AppConfig) {
	app = appConfig
}

func AddDefaultData(templData *models.TemplateData, request *http.Request) *models.TemplateData {
	templData.Flash = app.Session.PopString(request.Context(), "flash")
	templData.Error = app.Session.PopString(request.Context(), "errror")
	templData.Warning = app.Session.PopString(request.Context(), "warning")
	templData.CSRFToken = nosurf.Token(request)
	return templData
}

// RenderTemplate renders template using html/template
func RenderTemplate(writer http.ResponseWriter, request *http.Request, tmpl string, templData *models.TemplateData) error {
	var templateCache map[string]*template.Template

	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	templ, ok := templateCache[tmpl]
	if !ok {
		return errors.New("can't get template from cache")
	}

	buff := new(bytes.Buffer)

	templData = AddDefaultData(templData, request)

	_ = templ.Execute(buff, templData)

	_, err := buff.WriteTo(writer)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
		return err
	}

	return nil
}

// CreateTemplateCache creates a template cache as map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = templateSet
	}

	return myCache, nil
}
