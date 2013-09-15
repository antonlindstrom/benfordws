package main

import (
	"encoding/json"
	"net/http"
	"log"
	"fmt"
	"github.com/antonlindstrom/benfordslaw/counter"
	"io"
	)

type ErrorMessage struct {
	Error	string
}

type Dataset struct {
	Set []int
}

func JsonErrMsg(message string) string {
	var msg ErrorMessage = ErrorMessage{message}

	b, err := json.Marshal(msg)

	if err != nil {
		log.Printf("Failed to marshal JSON: %s\n", err)
	}

	return string(b[:])
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
		fmt.Fprintf(w, "%s\n", JsonErrMsg("Only supports method POST"))
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
