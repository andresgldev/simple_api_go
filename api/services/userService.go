package services

import (
	"github.com/andresgldev/simple_api_go/api/database"
	"github.com/andresgldev/simple_api_go/api/models"
)

type UserService struct {
}

func (c *UserService) All() ([]models.User, error) {
	var users = []models.User{}

	rows, err := database.Get().Query("select * from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (c *UserService) Show(id int) (*models.User, error) {
	var u models.User
	err := database.Get().QueryRow("select * from users where id = $1 limit 1", id).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		return nil, err
	}
	return u, nil
}

/*
func (c *UserService) Store(u *UserService) {
	json.NewDecoder(r.Body).Decode(&u)

	err := database.Get().QueryRow("insert into users (name, email) values ($1, $2) RETURNING id", u.Name, u.Email).Scan(&u.ID)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(u)
}
func (c *UserService) Update(w http.ResponseWriter, r *http.Request) {
	var u models.User
	json.NewDecoder(r.Body).Decode(&u)

	id := chi.URLParam(r, "id")

	_, err := database.Get().Exec("UPDATE users set name = $1, email = $2 where id = $3", u.Name, u.Email, id)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(u)
}
func (c *UserService) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete"))
}
*/
