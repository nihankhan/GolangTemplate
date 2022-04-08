package test

import (
	"testing"
	"net/http"
)

var client = &http.Client{}

func CheckRouter (t *testing.T) {
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080/", nil)

	if err != nil {
		t.Fail()
	}

	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != http.StatusOK {
		t.Fail()
	}	
}