package routes

import (
	"net/http"

	"github.com/andresgldev/simple_api_go/api/controllers"
	"github.com/andresgldev/simple_api_go/api/middlewares"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func InitRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi"))
	})

	r.Get("/prueba", controllers.Prueba)

	r.Route("/user", func(r chi.Router) {
		r.Use(middlewares.OnlyAdmin)
		r.Get("/", controllers.IndexUser)
	})

	return r
}
