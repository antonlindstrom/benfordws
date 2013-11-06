package main

import (
	"bytes"
	"io/ioutil"
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"regexp"
)


// Test default benford handler
func TestBenfordHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(BenfordHandler))
	defer server.Close()

	// Send GET request
	resp, err := http.Get(server.URL)

	if err != nil {
		t.Fatal("Error in http.Get!")
	}

	// Read Body
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	// Test if we get the Only supports method POST error on GET
	if string(body) != "{\"Error\":\"Only supports method POST\"}\n" {
		t.Fatal("Did not send error on GET request")
	}

	// Build a small test dataset
	var d Dataset
	d.Set = []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584}

	// Encode to JSON and post to server
	b, _ := json.Marshal(&d)
	resp, err = http.Post(server.URL, "application/json", bytes.NewReader(b))

	if err != nil {
		t.Fatal("Error in http.Post!")
	}

	// Read Body
	body, _ = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	re := regexp.MustCompile("30.10299956639812")

	// Test if we'll able to match the result with what we get from GET request
	if re.FindString(string(body)) != "30.10299956639812" {
		t.Fatal("Could not find value 30.10299956639812")
	}
}
