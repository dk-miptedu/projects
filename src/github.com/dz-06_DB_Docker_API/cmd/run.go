package cmd

import (
	"dz-06_DB_Docker_API/db"
	"dz-06_DB_Docker_API/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {
	db.Connect()
	db.Migrate()

	r := mux.NewRouter()

	r.HandleFunc("/transactions", handlers.HandleTransactions)
	r.HandleFunc("/transactions", handlers.Authenticate(handlers.HandleTransactions)).Methods("GET", "POST")
	r.HandleFunc("/transactions/{id}", handlers.HandleTransactions) // Для PUT и DELETE

	http.Handle("/", r)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
