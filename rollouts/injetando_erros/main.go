package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

var shouldBeGeneratingError string

func main() {
	http.HandleFunc("/", decideResult)
	fmt.Printf("GENERATE_ERROR: %s\n", shouldBeGeneratingError)
	fmt.Println("Starting Web server in :3000")
	http.ListenAndServe(":3000", nil)
}

func decideResult(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request received from %s", r.RemoteAddr)
	if shouldBeGeneratingError == "true" {
		if rand.Intn(2) == 0 {
			resultsOk(w, r)
		} else {
			resultsFail(w, r)
		}
	} else {
		resultsOk(w, r)
	}
}

func resultsOk(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
	log.Printf("Succeeded with code, %s", strconv.Itoa(http.StatusOK))
}

func resultsFail(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "Failed")
	log.Printf("Failed with code, %s", strconv.Itoa(http.StatusInternalServerError))
}
