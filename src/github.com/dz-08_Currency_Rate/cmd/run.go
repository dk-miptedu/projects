package cmd

import (
	"dz-08_Currency_Rate/db"
	"dz-08_Currency_Rate/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {
	fmt.Println("Starting the database...")
	//db.InitDB()
	fmt.Println("Connecting to the database...")
	db.Connect()
	fmt.Println("Connected to the database")
	fmt.Println("Migrating the database...")
	db.Migrate()

	r := mux.NewRouter()

	r.HandleFunc("/commissions/calculate", handlers.CalculateCommission)
	r.HandleFunc("/transactions", handlers.HandleTransactions)
	r.HandleFunc("/transactions", handlers.Authenticate(handlers.HandleTransactions)).Methods("GET", "POST")
	r.HandleFunc("/transactions/{id}", handlers.HandleTransactions) // Для PUT и DELETE
	r.HandleFunc("/users", handlers.RegisterUser).Methods("POST")
	r.HandleFunc("/users/login", handlers.LoginUser).Methods("POST")

	http.Handle("/", r)

	fmt.Println("Server is running on port 8080")
	fmt.Println("Press Ctrl+C to quit.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
