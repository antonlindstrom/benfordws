package main

import (
	"encoding/json"
	"net/http"
	"log"
	"fmt"
	"github.com/antonlindstrom/benfordslaw/counter"
	"io"
	)

type Response map[string]interface{}

type Dataset struct {
	Set []int
}

func (r Response) String() (string) {
	b, err := json.Marshal(r)

	if err != nil {
		log.Printf("Failed to marshal JSON: %s\n", err)
		return ""
	}

	return string(b)
}

func parseSet(body io.ReadCloser) ([]int) {
	var d Dataset

	decoder := json.NewDecoder(body)
	err     := decoder.Decode(&d)

	if err != nil {
		log.Printf("Error could not parse JSON!\n")
	}

	return d.Set
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "%s\n", Response{"Error":"Only supports method POST"})
		return
	}

	total, _ := counter.Process(parseSet(r.Body))
	log.Printf("Processed dataset of %d numbers\n", total)
}

func main() {
	log.Printf("Starting up WS on port 6000")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":6000", nil))
}
