package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myH myHandler
	handler := NoSurf(&myH)

	switch handler.(type) {
	case http.Handler:
		//Test pass - type is correct
	default:
		t.Error("Type is not http.Handler")
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler
	handler := SessionLoad(&myH)

	switch handler.(type) {
	case http.Handler:
		//Test pass - type is correct
	default:
		t.Error("Type is not http.Handler")
	}
}
