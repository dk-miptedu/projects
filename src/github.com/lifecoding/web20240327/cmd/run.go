package cmd

import (
	"database/sql"
	"lessons/handlers"
	"log"
	"net/http"
)

func Run(db *sql.DB) {
	http.HandleFunc("/item", handlers.Item(db))
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
