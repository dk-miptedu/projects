package cmd

import (
	"dz-06_DB_Docker_API/handlers"
	"dz-06_DB_Docker_API/repo"
	"fmt"
	"net/http"
)

func Run() {
	repo.InitDB()
	fmt.Println("Web server starts...")

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/greet", handlers.GreetingHandler)
	mux.HandleFunc("/json", handlers.JsonHandler)
	mux.HandleFunc("/transactions", handlers.TransactionsHandler)
	mux.HandleFunc("/transactions/", handlers.TransactionsHandler)

	loggedMux := handlers.LoggingMiddleware(mux)

	http.ListenAndServe(":8080", loggedMux)

}
