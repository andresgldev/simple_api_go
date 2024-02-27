package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func Run() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	fmt.Println(os.Getenv("DB_USER"))

	fmt.Println("RUN")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})
	http.HandleFunc("/prueba", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Lugar de prueba")
	})

	http.ListenAndServe(":8080", nil)

	fmt.Println("FIN")
}
