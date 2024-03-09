package handlers

import (
	"dz-05_RESTful_API/models"
	"dz-05_RESTful_API/repo"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var InMemoryDB = repo.NewInMemoryDB()

//InMemoryDB := repo.InMemoryDB

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go Web Server! Now: %v", time.Now())
}

func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s! Now: %v", name, time.Now())
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed in %v", time.Since(start))
	})
}

// Обработчик для parh - ./json
// curl --header "Content-Type: application/json" --request POST --data '{"name":"theName","password":"it password attr"}' http://localhost:8080/json
func JsonHandler(w http.ResponseWriter, r *http.Request) {
	var userReq models.UserRequest
	// Десериализация JSON из тела запроса в структуру
	json.NewDecoder(r.Body).Decode(&userReq)
	// Сериализация и отправка JSON ответа
	userResp := models.UserResponse{Greeting: "Hello, " + userReq.Name + ", use passwd:" + userReq.Password}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userResp)
}

// curl --header "Content-Type: application/json" --request POST --data '{"id":"01","amount":120.30,"currency":"USD","type":"income","category":"Зарплата","date":"2024-03-01","description":"Описание"}' http://localhost:8080/transactions
// GET POST
func TransactionsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":

		var item models.Transaction
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		currentId := InMemoryDB.Create(item)
		// Сериализация и отправка JSON ответа
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.ItemResponse{Item: currentId, Ok: "Ok"})

	case "GET":
		//var items []models.Transaction
		w.Header().Set("Content-Type", "application/json")
		//path := strings.Split(r.URL.Path, "/")[2]
		//fmt.Println("Provision ID: ", path)
		//json.NewEncoder(w).Encode(items)
		json.NewEncoder(w).Encode(&InMemoryDB.Items)
	case "PUT":
	case "DELETE":
	}

}
