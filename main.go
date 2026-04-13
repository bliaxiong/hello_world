// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintln(w, "Hello World"); err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// Dummy function for testing in pipeline
func strLen(str string) int {
	return len(str)
}
