package api

import (
	"fmt"
	"net/http"

	"github.com/andresgldev/simple_api_go/api/config"
	"github.com/andresgldev/simple_api_go/api/database/migration"
	"github.com/andresgldev/simple_api_go/api/routes"
)

func Run() {

	config := config.Get()
	migration.Run()
	routes := routes.InitRoutes()
	fmt.Printf("Service run in :%d", config.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), routes)
}
