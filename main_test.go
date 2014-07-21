package main

import (
	"net/http/httptest"
	"testing"
)

func TestHandle(t *testing.T) {
	rec := httptest.NewRecorder()
	handle(rec, nil)
	expected := "Hello, World!"
	if body := rec.Body.String(); body != expected {
		t.Errorf("Expected: [%s] Recieved: [%s]", expected, body)
	}
}
