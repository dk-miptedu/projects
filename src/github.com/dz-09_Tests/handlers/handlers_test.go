package handlers_test

import (
	"bytes"
	"encoding/json"
	"golang_hws/handlers"
	"golang_hws/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

func TestHandleTransactionsGet(t *testing.T) {
	// Setup
	r := mux.NewRouter()
	r.HandleFunc("/transactions", handlers.HandleTransactions).Methods("GET")
	ts := httptest.NewServer(r)
	defer ts.Close()

	// Execute
	res, err := http.Get(ts.URL + "/transactions")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	defer res.Body.Close()

	// Verify
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}

	// Here you'd check the response body for correct data
	// This step is omitted for simplicity
}

func TestHandleTransactionsPost(t *testing.T) {
	// Setup
	r := mux.NewRouter()
	r.HandleFunc("/transactions", handlers.HandleTransactions).Methods("POST")
	ts := httptest.NewServer(r)
	defer ts.Close()
	dateTimeStr := "2023-03-13T14:00:00Z"
	parsedTime, err := time.Parse(time.RFC3339, dateTimeStr)
	// Create a transaction to send in the request
	newTransaction := []models.Transactions{
		{ID: 1, UserID: 1, Amount: 100.00, Currency: "USD", TransactionType: "transfer", Category: "SBP", TransactionDate: parsedTime, Description: "Service 01", Commission: 1.00},
		{ID: 2, UserID: 1, Amount: 200.00, Currency: "EUR", TransactionType: "transfer", Category: "SBP", TransactionDate: parsedTime, Description: "Service 02", Commission: 5.00},
	}

	body, _ := json.Marshal(newTransaction)

	// Execute
	res, err := http.Post(ts.URL+"/transactions", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("Failed to send POST request: %v", err)
	}
	defer res.Body.Close()

	// Verify
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, res.StatusCode)
	}

	// You would also typically unmarshal and check the response body here
	// This step is omitted for simplicity
}
