package models

import (
	"strings"
	"time"
)

type Users struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Username     string         `json:"name"`
	Email        string         `gorm:"unique" json:"email"`
	Password     string         `json:"password"`
	Transactions []Transactions `gorm:"foreignKey:UserID"`
}

type Transactions struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	UserID          uint      `gorm:"not null" json:"user_id"`
	Amount          float64   `gorm:"not null;default:0" json:"amount"`
	Currency        string    `gorm:"not null" json:"currency"`
	TransactionType string    `json:"type"`
	Category        string    `json:"category"`
	TransactionDate time.Time `json"json:"date"`
	Description     string    `json:"description"`
	Commission      float64   `gorm:"not null;default:0" json:"commission"`
	//User            *Users    `gorm:"foreignKey:UserID"`
}

// Проверка: Ограничение по типу транзакции - "income", "expense" или "transfer"
func IsTransactionType(s string) bool {
	switch strings.TrimSpace(s) {
	case "income", "expense", "transfer":
		return true

	}
	return false

}
