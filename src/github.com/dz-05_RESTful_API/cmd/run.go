package cmd

import (
	"dz-05_RESTful_API/handlers"
	"fmt"
	"net/http"
)

func Run() {
	fmt.Println("Web server starts...")
	//http.HandleFunc("/", handlers.HomeHandler)
	//http.HandleFunc("/greet", handlers.GreetingHandler)

	//http.ListenAndServe(":8090", nil)

	mux := http.NewServeMux()
	//mux.HandleFunc("/", handlers.)

	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/greet", handlers.GreetingHandler)
	mux.HandleFunc("/json", handlers.JsonHandler)
	mux.HandleFunc("/transactions", handlers.TransactionsHandler)
	mux.HandleFunc("/transactions/{id}", handlers.TransactionsHandler)
	loggedMux := handlers.LoggingMiddleware(mux)
	http.ListenAndServe(":8080", loggedMux)

}
