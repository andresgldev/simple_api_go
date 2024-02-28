package routes

import (
	"database/sql"
	"net/http"

	"github.com/andresgldev/simple_api_go/api/controllers"
	"github.com/andresgldev/simple_api_go/api/middlewares"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func InitRoutes(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", (controllers.HomeController{}).Index)
	r.Get("/demo", (controllers.HomeController{}).Demo)

	r.Route("/user", func(r chi.Router) {
		u := controllers.UserController{DB: db}
		r.Use(middlewares.OnlyAdmin)
		r.Get("/", u.Index)
		r.Post("/", u.Store)
		r.Get("/{id}", u.Show)
		r.Put("/{id}", u.Update)
		r.Delete("/{id}", u.Delete)
	})

	/*r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404", http.StatusNotFound)
	})*/

	return r
}
