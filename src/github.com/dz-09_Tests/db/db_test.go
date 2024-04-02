package db

import (
	"golang_hws/models"
	"testing"
	"time"
)

func TestUsersTransactionsIntegration(t *testing.T) {
	Connect()
	Migrate()
	defer Close() //DB.Migrator().DropTable(&models.Users{}, &models.Transactions{})

	// Create a user
	user := models.Users{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password",
	}
	if err := DB.Create(&user).Error; err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// Create a transaction
	transaction := models.Transactions{
		UserID:          user.ID,
		Amount:          100.0,
		Currency:        "USD",
		TransactionType: "income",
		Category:        "Test",
		TransactionDate: time.Now(),
		Description:     "Test Transaction",
		Commission:      2.0,
	}
	if err := DB.Create(&transaction).Error; err != nil {
		t.Fatalf("Failed to create transaction: %v", err)
	}

	// Fetch and verify the transaction
	var fetchedTx models.Transactions
	if err := DB.First(&fetchedTx, transaction.ID).Error; err != nil {
		t.Fatalf("Failed to fetch transaction: %v", err)
	}

	if !models.IsTransactionType(fetchedTx.TransactionType) {
		t.Errorf("Invalid transaction type: %v", fetchedTx.TransactionType)
	}
}
