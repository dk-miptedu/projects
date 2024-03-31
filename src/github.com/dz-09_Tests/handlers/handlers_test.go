package handlers

import (
	"encoding/json"
	"golang_hws/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestGetTransactions tests the getTransactions function
func TestGetTransactions(t *testing.T) {
	// Initialize a new httptest.ResponseRecorder
	rr := httptest.NewRecorder()

	// Initialize a new http request
	req, err := http.NewRequest("GET", "/transactions", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Assuming a global variable or a way to inject a mock db instance exists
	// Mock the expected response
	dateTimeStr := "2023-03-13T14:00:00Z"
	parsedTime, err := time.Parse(time.RFC3339, dateTimeStr)
	mockedTransactions := []models.Transactions{
		{ID: 1, UserID: 1, Amount: 100.00, Currency: "USD", TransactionType: "transfer", Category: "SBP", TransactionDate: parsedTime, Description: "Service 01", Commission: 1.00},
		{ID: 2, UserID: 1, Amount: 200.00, Currency: "EUR", TransactionType: "transfer", Category: "SBP", TransactionDate: parsedTime, Description: "Service 02", Commission: 5.00},
	}

	getTransactions(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var transactions []models.Transactions
	err = json.NewDecoder(rr.Body).Decode(&transactions)
	if err != nil {
		t.Fatal("Failed to decode response body")
	}

	if len(transactions) != 2 {
		t.Errorf("handler returned unexpected number of transactions: got %v want %v", len(transactions), 2)
	}
}
