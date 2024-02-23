package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Привет! Добро пожаловать в BudgetCalcSimple!")
	budget := Transactions{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Добавьте запись в формате: <Категория:Сумма>")
	fmt.Println("Доход это Сумма со знаком плюс, Расход это Сумма со знаком минус")
	i := 1
	TransactionsNum(i)
	for scanner.Scan() {
		rowAdd := scanner.Text()
		if rowAdd == "end" {
			fmt.Println("Транзакции закончились, начинаем расчет...")
			break
		}

		transaction := strings.Split(rowAdd, ":")
		if len(transaction) != 2 {
			fmt.Printf("Ошибка при распознавании транзакции\n Вы ввели %s\n Транзакция не принимается, попробуйте еще раз\n", rowAdd)
			fmt.Println("попробуйте еще раз")
			TransactionsNum(i)
			continue
		}
		amount, err := strconv.ParseFloat(transaction[1], 64)
		if err != nil {
			fmt.Printf("Ошибка при распознавании Суммы транзакции\n Вы ввели %s\n Транзакция не принимается, попробуйте еще раз\n", rowAdd)
			fmt.Println("попробуйте еще раз")
			TransactionsNum(i)
			continue
		}
		budget.AddTransaction(transaction[0], amount)
		i += 1
		TransactionsNum(i)
	}

	balance, income := budget.CalcBalanceAndIncome()
	fmt.Printf("Чистый доход: %.2f\n", balance)
	budget.CalcAnalyze(income)

}
