package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	const PORT = "4000"
	fmt.Printf("server @ localhost:%s\n", PORT)
	err = http.ListenAndServe(":"+PORT, mux)
	log.Fatal(err)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Printf("%s %s\n", r.Method, r.URL.Path)

	// Common code for all requests...

	switch r.Method {
	case http.MethodGet:
		// Handle the GET request...
		fmt.Fprintf(w, "handling GET request\n")

	case http.MethodOptions:
		w.Header().Set("Allow", "GET, OPTIONS")
		w.WriteHeader(http.StatusNoContent)

	default:
		w.Header().Set("Allow", "GET, OPTIONS")
		http.Error(w, "error: method not allowed", http.StatusMethodNotAllowed)
	}
}
