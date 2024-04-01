package handlers

import (
	"golang_hws/models"
	"testing"
	"time"
)

func TestaddTransaction(t *testing.T) {
	// Setup
	parsedTime := time.Now()
	// Тестовые данные, которые ожидаем добавить в Базу данных
	testData := []models.Transactions{
		{ID: 1, UserID: 1, Amount: 100.00, Currency: "USD", TransactionType: "transfer", Category: "SBP", TransactionDate: parsedTime, Description: "Service 01", Commission: 2.00},
		{ID: 2, UserID: 1, Amount: 200.00, Currency: "EUR", TransactionType: "transfer", Category: "SBP", TransactionDate: parsedTime, Description: "Service 02", Commission: 4.00},
		{ID: 3, UserID: 1, Amount: 1000.00, Currency: "GBP", TransactionType: "transfer", Category: "SBP", TransactionDate: parsedTime, Description: "Service 02", Commission: 20.0},
	}
	// Execute & Verify
	//Проверяем Расчет комиссии и обновление щтампа времени транзакции
	for _, aa := range testData {
		a := aa.TransactionDate //Запоминаем штамп времени из testData
		b := aa.Commission      //Запоминаем размер комиссии из testData
		CalculateCommission(&aa)
		aa.TransactionDate = time.Now()
		if b != aa.Commission || a.After(aa.TransactionDate) {
			t.Errorf("Expected Comission: %f, got %f", b, aa.Commission)
			t.Errorf("Expected TransactionDate: %v, got %v", a, aa.TransactionDate)
			continue
		}
	}
}
