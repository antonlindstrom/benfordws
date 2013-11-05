package main

import (
	"encoding/json"
	"fmt"
	"github.com/antonlindstrom/benfordslaw"
	"github.com/technoweenie/grohl"
	"log"
	"net/http"
	"os"
)

type Dataset struct {
	Set []int
}

// Main, start up WS
func main() {
	var port = "6000"

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	grohl.AddContext("app", "benfordws")
	grohl.Log(grohl.Data{"state": "startup", "port": port})

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// Handler to process the data. Only Allows POST
func handler(w http.ResponseWriter, r *http.Request) {
	timer := grohl.NewTimer(grohl.Data{"path": r.URL.Path, "method": r.Method})

	defer timer.Finish()
	defer r.Body.Close()

	if r.Method != "POST" {
		fmt.Fprintf(w, "%s\n", "{\"Error\":\"Only supports method POST\"}")
		return
	}

	var d Dataset

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&d)

	if err != nil {
		grohl.Log(grohl.Data{"error": "json", "type": "decode"})
	}

	fmt.Fprintf(w, "%s\n", benfordslaw.Process(d.Set))
}
