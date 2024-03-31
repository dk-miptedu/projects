package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang_hws/db"
	"golang_hws/models"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
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
	var t models.Transactions
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	CalculateCommission(&t)
	t.TransactionDate = time.Now()

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
	var transactions []models.Transactions
	userid, _ := ParseTokenFromRequest(r)
	//fmt.Println("userid: ", userid)
	if userid == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}
	result := db.DB.Where("user_id = ?", userid).Find(&transactions)
	//result := db.DB.Find(&transactions)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactions)
}

// updateTransaction обновляет данные транзакции в базе данных
func updateTransaction(w http.ResponseWriter, r *http.Request) {
	var t models.Transactions
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	userid, _ := ParseTokenFromRequest(r)
	if userid == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}
	if err := db.DB.Where("user_id = ? AND id = ?", userid, id).First(&t).Error; err != nil {
		http.Error(w, "Transactions user_id not found", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t.ID = uint(id)
	CalculateCommission(&t)
	t.TransactionDate = time.Now()
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
	transactionDeleted := "Transaction deleted"
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	userid, _ := ParseTokenFromRequest(r)
	//fmt.Println("userid: ", userid)
	if userid == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	result := db.DB.Where("user_id = ?", userid).Delete(&models.Transactions{}, id)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	} else if result.RowsAffected == 0 {
		// No rows were deleted, which might be unexpected
		log.Printf("No records found to delete for user_id: %v with id: %v", userid, id)
		transactionDeleted = "No User's records found to delete"
	} else {
		log.Printf("Successfully deleted record.")
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(transactionDeleted))
}
func CalculateCommission(c *models.Transactions) {

	switch {
	case c.TransactionType == "transfer" && c.Currency == "USD", c.Currency == "EUR", c.Currency == "GBP", c.Currency == "JPY":
		c.Commission = c.Amount * 0.02
	case c.TransactionType == "transfer" && c.Currency == "RUB":
		c.Commission = c.Amount * 0.05
	case c.TransactionType == "purchase", c.TransactionType == "top-up":
		c.Commission = 0
	default:
		c.Commission = 0
	}

}
func ParseTokenFromRequest(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header is required")
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("YourSigningKey"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userID := fmt.Sprintf("%.0f", claims["user_id"].(float64))
		//fmt.Printf("claims: %s\n", claims)
		return userID, nil
	} else {
		return "", err
	}
}
