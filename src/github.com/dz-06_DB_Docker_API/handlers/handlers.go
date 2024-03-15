package handlers

import (
	"dz-06_DB_Docker_API/db"
	"dz-06_DB_Docker_API/models"
	"encoding/json"
	"net/http"
	"strconv"

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
