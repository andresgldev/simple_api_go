package migration

import (
	"log"

	"github.com/andresgldev/simple_api_go/api/database"
)

func Run() {
	db := database.Get()

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")
	if err != nil {
		log.Fatal(err)
	}
}
