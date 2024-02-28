package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserController struct {
	DB *sql.DB
}

func (c UserController) Index(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Index"))

	rows, err := c.DB.Query("select * from users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var u User
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
	var u User
	err := c.DB.QueryRow("select * from users where id = $1 limit 1", id).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(u)
}
func (c UserController) Store(w http.ResponseWriter, r *http.Request) {
	var u User
	json.NewDecoder(r.Body).Decode(&u)

	err := c.DB.QueryRow("insert into users (name, email) values ($1, $2) RETURNING id", u.Name, u.Email).Scan(&u.ID)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(u)
}
func (c UserController) Update(w http.ResponseWriter, r *http.Request) {
	var u User
	json.NewDecoder(r.Body).Decode(&u)

	id := chi.URLParam(r, "id")

	_, err := c.DB.Exec("UPDATE users set name = $1, email = $2 where id = $3", u.Name, u.Email, id)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(u)
}
func (c UserController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete"))
}
