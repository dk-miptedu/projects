package delete0

import (
	"fmt"
	"golang_hws/handlers"
	"golang_hws/models"
	"time"
)

func TestAddTrn() {
	//dateTimeStr := "2023-03-13T14:00:00Z"
	//parsedTime, err := time.Parse(time.RFC3339, dateTimeStr)
	parsedTime := time.Now() //Parse(time.RFC3339, dateTimeStr)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	// TODO: Implement
	t := []models.Transactions{
		{ID: 1, UserID: 1, Amount: 100.00, Currency: "USD", TransactionType: "transfer", Category: "SBP", TransactionDate: parsedTime, Description: "Service 01", Commission: 2.00},
		{ID: 2, UserID: 1, Amount: 200.00, Currency: "EUR", TransactionType: "transfer", Category: "SBP", TransactionDate: parsedTime, Description: "Service 02", Commission: 4.00},
		{ID: 3, UserID: 1, Amount: 1000.00, Currency: "GBP", TransactionType: "transfer", Category: "SBP", TransactionDate: parsedTime, Description: "Service 02", Commission: 20.0},
	}
	//CalculateCommission(&t)
	//t.TransactionDate = time.Now()
	for in, aa := range t {
		//fmt.Println(aa.TransactionDate)
		a := aa.TransactionDate
		b := aa.Commission
		handlers.CalculateCommission(&aa)
		aa.TransactionDate = time.Now()
		if b == aa.Commission && a.Before(aa.TransactionDate) {
			fmt.Println(in, " ", aa)
			continue
		}

		//fmt.Println(in, " ", aa)

	}
}
