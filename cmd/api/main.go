package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})
	http.HandleFunc("/prueba", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Lugar de prueba")
	})
	http.ListenAndServe(":8080", nil)
}
