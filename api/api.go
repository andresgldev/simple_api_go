package api

import (
	"net/http"

	"github.com/andresgldev/simple_api_go/api/routes"
	"github.com/joho/godotenv"
)

func Run() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	// fmt.Println(os.Getenv("DB_USER"))

	// fmt.Println("RUN")

	routes := routes.InitRoutes()
	http.ListenAndServe(":8080", routes)
}
