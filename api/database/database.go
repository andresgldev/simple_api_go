package database

import (
	"database/sql"
	"fmt"

	"github.com/andresgldev/simple_api_go/api/config"
	_ "github.com/lib/pq"
)

var database *sql.DB

func Get() *sql.DB {
	if database == nil {
		database = initDatabase()
	}
	return database
}

func initDatabase() *sql.DB {
	c := config.Get()
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", c.DBhost, c.DBuser, c.DBpass, c.DBname))
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	return db
}
