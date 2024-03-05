package cmd

import (
	"log"
	"net/http"
	"web20240301/handlers"
)

func Run() {

	http.HandleFunc("/item", handlers.Item)
	log.Println("server started on: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
