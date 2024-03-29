package models

import (
	"strings"

	"github.com/rmg/iso4217"
)

type UserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserResponse struct {
	Greeting string `json:"greeting"`
}

type ItemResponse struct {
	Item string `json:"id"`
	Ok   string `json:"status"`
}

type Transaction struct {
	ID          string  `json:"id"`
	Amount      float32 `json:"amount"`      //сумма
	Currency    string  `json:"currency"`    //валюта: "USD", "EUR", "RUB", "JPY", и т. д.
	Types       string  `json:"type"`        //тип: доход, расход или перевод
	Category    string  `json:"category"`    //категория: зарплата, еда, жильё и т. д.
	Dates       string  `json:"date"`        //дата
	Description string  `json:"description"` //описание

}
type TransactionList struct {
	Item []Transaction `json:"item"`
	Ok   bool          `json:"Ok"`
}

// 202024-03-09 Проверки в проекте не используются, добавлю потом
//
// Проверка: Ограничение по типу транзакции - "income", "expense" или "transfer"
func IsTransactionType(s string) bool {
	switch strings.TrimSpace(s) {
	case "income", "expense", "transfer":
		return true

	}
	return false

}

// Проверка: Существует ли Код валюты в справочнике ISO
func IsCurrency(cur string) bool {
	code, _ := iso4217.ByName(strings.TrimSpace(cur))
	if code == 0 {
		return false
	}
	return true
}
