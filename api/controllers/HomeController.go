package controllers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my website!")
}

func Prueba(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Lugar de prrre")
}
