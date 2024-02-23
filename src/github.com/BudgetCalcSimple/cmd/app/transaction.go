package main

import (
	"fmt"
	"math"
)

type Transactions struct {
	Amount map[string]float64
}

// Функция для добавления записи Доход/Расход по категориям
func (f *Transactions) AddTransaction(category string, amount float64) {
	if f.Amount == nil {
		f.Amount = make(map[string]float64)
	}
	f.Amount[category] += amount
}

// Функция для расчета Баланса - Доходов и Расходов
func (f *Transactions) CalcBalanceAndIncome() (float64, float64) {
	totalAmount := 0.0
	income := 0.0
	for _, amount := range f.Amount {
		totalAmount += amount
		if amount > 0.0 {
			income += amount
		}
	}
	return totalAmount, income
}

// Функция для анализа движения Денежных средств по категориям
func (f *Transactions) CalcAnalyze(income float64) {
	for category, amount := range f.Amount {
		attention := [3]string{"", "Внимание:", "Рекомендуется снизить траты в этой категории."}

		percentage := math.Abs(amount/income) * 100
		if amount <= 0.0 {
			if math.Ceil(percentage) >= 30 {
				fmt.Printf("%sРасходы на %s составляют %.2f%% от вашего дохода. %s\n", attention[1], category, percentage, attention[2])
			} else {
				fmt.Printf("%sРасходы на %s составляют %.2f%% от вашего дохода.\n", attention[0], category, percentage)
			}
		} else {
			fmt.Printf("Доходы по %s составляют %.2f%% от поступлений\n", category, percentage)
		}
	}
}

// Функция отображения порядкового номера Транзакции для ввода
func TransactionsNum(i int) {
	fmt.Printf("Транзакция %d:>", i)
}
