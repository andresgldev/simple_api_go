package controllers

import (
	"net/http"
)

type HomeController struct {
}

func (c HomeController) Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi"))
}

func (c HomeController) Demo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Demo"))
}
