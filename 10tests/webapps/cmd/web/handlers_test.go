package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_handlers(t *testing.T) {
	theTests := []struct {
		name               string
		url                string
		expectedStatusCode int
	}{
		{"home", "/home", http.StatusOK},
		{"404", "/not_exist", http.StatusNotFound},
	}

	var app application
	routes := app.routes()

	//create a test server
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	pathToTemplates = "./../../templates/"

	for _, e := range theTests {
		resp, err := ts.Client().Get(ts.URL + e.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}
		if resp.StatusCode != e.expectedStatusCode {
			t.Errorf("for %s: expected status %d, but got  %d", e.name, e.expectedStatusCode, resp.StatusCode)
		}

	}
}
