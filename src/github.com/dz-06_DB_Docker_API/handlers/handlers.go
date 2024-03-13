package handlers

import (
	"dz-06_DB_Docker_API/models"
	"dz-06_DB_Docker_API/repo"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	statusSuccess = "Success"
	statusError   = "Error"
)

var InMemoryDB = repo.NewInMemoryDB()

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
	json.NewDecoder(r.Body).Decode(&userReq)
	userResp := models.UserResponse{Greeting: "Hello, " + userReq.Name + ", use passwd:" + userReq.Password}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userResp)
}

// curl --header "Content-Type: application/json" --request POST --data '{"id":"01","amount":120.30,"currency":"USD","type":"income","category":"Зарплата","date":"2024-03-01","description":"Описание"}' http://localhost:8080/transactions
// GET POST
func TransactionsHandler(w http.ResponseWriter, r *http.Request) {
	var item models.Transaction
	id := strings.TrimPrefix(r.URL.Path, "/transactions/")
	idInt, _ := strconv.Atoi(id)
	fmt.Println("Transactions id: %n", idInt, " Path", r.URL.Path)

	switch r.Method {
	case "POST":

		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.Header().Set("Content-Type", "application/json")
		currentId := InMemoryDB.CreateItem(item)
		if len(currentId) > 0 {
			// Сериализация и отправка JSON ответа
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(models.ItemResponse{Item: currentId, Ok: statusSuccess})
		} else {
			json.NewEncoder(w).Encode(models.ItemResponse{Item: "", Ok: statusError})
		}

	case "GET":
		if idInt > 0 {
			for _, items := range InMemoryDB.Items {
				if items.ID == id {
					json.NewEncoder(w).Encode(items)
					return
				}
			}
			json.NewEncoder(w).Encode(models.ItemResponse{Item: "[]", Ok: statusSuccess})
			return
		} else if idInt < 0 {
			json.NewEncoder(w).Encode(models.ItemResponse{Item: "[]", Ok: statusSuccess})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&InMemoryDB.Items)
	case "PUT":
		w.Header().Set("Content-Type", "application/json")
		if idInt > 0 {
			_ = json.NewDecoder(r.Body).Decode(&item)
			//fmt.Println(id, ": ", item)
			status := InMemoryDB.UpdateItem(id, item)
			if status {
				w.WriteHeader(http.StatusAccepted)
				json.NewEncoder(w).Encode(models.ItemResponse{Item: id, Ok: statusSuccess})
			} else {
				w.WriteHeader(http.StatusConflict)
				json.NewEncoder(w).Encode(models.ItemResponse{Item: id, Ok: statusError})
			}
		} else {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(models.ItemResponse{Item: id, Ok: statusError})

		}

	case "DELETE":
		w.Header().Set("Content-Type", "application/json")
		if idInt > 0 {
			status := InMemoryDB.DeleteItem(id)
			if status {
				w.WriteHeader(http.StatusAccepted)
				json.NewEncoder(w).Encode(models.ItemResponse{Item: id, Ok: statusSuccess})
			} else {
				w.WriteHeader(http.StatusConflict)
				json.NewEncoder(w).Encode(models.ItemResponse{Item: id, Ok: statusError})
			}
		} else {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(models.ItemResponse{Item: id, Ok: statusError})

		}

	}

}
