package render

import (
	"net/http"
	"testing"

	"github.com/sseudes108/Course_Go_Web_App/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	var templData models.TemplateData

	request, err := getSession()
	if err != nil {
		t.Error(err)
	}

	session.Put(request.Context(), "flash", "123")

	result := AddDefaultData(&templData, request)

	if result.Flash != "123" {
		t.Error("flash value of 123 not found in session")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"

	request, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var writer myWriter

	err = RenderTemplate(&writer, request, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error("Error writing template to browser")
	}

	err = RenderTemplate(&writer, request, "non-existent.page.tmpl", &models.TemplateData{})
	if err == nil {
		t.Error("Error renderer template that do not exist")
	}

}

func TestNewTemplates(t *testing.T) {
	NewTemplates(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"

	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}

func getSession() (*http.Request, error) {
	request, err := http.NewRequest("GET", "/url-test", nil)
	if err != nil {
		return nil, err
	}

	context := request.Context()
	context, _ = session.Load(context, request.Header.Get("X-Session"))
	request = request.WithContext(context)

	return request, nil
}
