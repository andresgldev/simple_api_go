package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andresgldev/simple_api_go/api/database"
	"github.com/andresgldev/simple_api_go/api/models"
	"github.com/go-chi/chi"
)

type UserController struct {
}

func (c UserController) Index(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Index"))

	rows, err := database.Get().Query("select * from users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(users)
}
func (c UserController) Show(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var u models.User
	err := database.Get().QueryRow("select * from users where id = $1 limit 1", id).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(u)
}
func (c UserController) Store(w http.ResponseWriter, r *http.Request) {
	var u models.User
	json.NewDecoder(r.Body).Decode(&u)

	err := database.Get().QueryRow("insert into users (name, email) values ($1, $2) RETURNING id", u.Name, u.Email).Scan(&u.ID)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(u)
}
func (c UserController) Update(w http.ResponseWriter, r *http.Request) {
	var u models.User
	json.NewDecoder(r.Body).Decode(&u)

	id := chi.URLParam(r, "id")

	_, err := database.Get().Exec("UPDATE users set name = $1, email = $2 where id = $3", u.Name, u.Email, id)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(u)
}
func (c UserController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete"))
}
