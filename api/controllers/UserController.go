package controllers

import "net/http"

func IndexUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("asas"))
}
