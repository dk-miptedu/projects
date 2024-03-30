package handlers

import (
	"encoding/json"
	"golang_hws/db"
	"golang_hws/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

const (
	statusSuccess = "Success"
	statusError   = "Error"
)

// HandleTransactions роутер для обработки запросов к транзакциям
func HandleTransactions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getTransactions(w, r)
	case "POST":
		addTransaction(w, r)
	case "PUT":
		updateTransaction(w, r)
	case "DELETE":
		deleteTransaction(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Method not supported"))
	}
}

// addTransaction добавляет новую транзакцию в базу данных
func addTransaction(w http.ResponseWriter, r *http.Request) {
	var t models.Transaction
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	CalculateCommission(&t)
	t.Date = time.Now()

	result := db.DB.Create(&t)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}

// getTransactions возвращает список всех транзакций из базы данных
func getTransactions(w http.ResponseWriter, r *http.Request) {
	var transactions []models.Transaction
	result := db.DB.Find(&transactions)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactions)
}

// updateTransaction обновляет данные транзакции в базе данных
func updateTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var t models.Transaction
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t.ID = uint(id)
	result := db.DB.Save(&t)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(t)
}

// deleteTransaction удаляет транзакцию из базы данных
func deleteTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	result := db.DB.Delete(&models.Transaction{}, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Transaction deleted"))
}
func CalculateCommission(c *models.Transaction) {

	switch {
	case c.Type == "transfer" && c.Currency == "USD":
		c.Commission = c.Amount * 0.02
	case c.Type == "transfer" && c.Currency == "RUB":
		c.Commission = c.Amount * 0.05
	case c.Type == "purchase", c.Type == "top-up":
		c.Commission = 0
	default:
		c.Commission = 0
	}

}
