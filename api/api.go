package api

import (
	"fmt"
	"net/http"

	"github.com/andresgldev/simple_api_go/api/config"
	"github.com/andresgldev/simple_api_go/api/routes"
	_ "github.com/lib/pq"
)

func Run() {

	config := config.Get()
	//db := database.Get()

	/*_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")
	if err != nil {
		log.Fatal(err)
	}*/

	routes := routes.InitRoutes()
	fmt.Printf("Service run in :%d", config.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), routes)
}
