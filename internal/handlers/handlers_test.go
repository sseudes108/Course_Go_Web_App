package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"generals", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"majors", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"search-availability", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"make-reservation", "/make-reservation", "GET", []postData{}, http.StatusOK},

	{"post-search", "/search-availability", "POST", []postData{
		{key: "start", value: "2024-01-01"},
		{key: "end", value: "2024-02-02"},
	}, http.StatusOK},

	{"post-availability-json", "/search-availability-json", "POST", []postData{
		{key: "start", value: "2024-01-01"},
		{key: "end", value: "2024-02-02"},
	}, http.StatusOK},

	{"post-make-reservation", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "Alessandra"},
		{key: "last_name", value: "Ribeiro"},
		{key: "email", value: "Ale@she.io"},
		{key: "phone", value: "551196979895969"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	testServer := httptest.NewTLSServer(routes)
	defer testServer.Close()

	for _, test := range theTests {
		if test.method == "GET" {
			response, err := testServer.Client().Get(testServer.URL + test.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if response.StatusCode != test.expectedStatusCode {
				t.Errorf("For %s, expected %d but got %d", test.name, test.expectedStatusCode, response.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, currentTest := range test.params {
				values.Add(currentTest.key, currentTest.value)
			}
			response, err := testServer.Client().PostForm(testServer.URL+test.url, values)

			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if response.StatusCode != test.expectedStatusCode {
				t.Errorf("For %s, expected %d but got %d", test.name, test.expectedStatusCode, response.StatusCode)
			}
		}
	}
}
