package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/andresgldev/simple_api_go/api/routes"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Run() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME")))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Conectado a la bd")
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	routes := routes.InitRoutes(db)
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), routes)
}
