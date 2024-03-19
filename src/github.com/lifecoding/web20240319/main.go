package main

import "fmt"

func main() {
	// TODO: implement main function
	fmt.Println("Для каждого проекта вычисляем долю от общего бюджета исходя из его приоритета")
	projects := []Project{
		{"Проект A", 5},
		{"Проект B", 3},
		{"Проект C", 2},
	}
	totalBudget := 1000
	fmt.Println("\nПриоритеты проектов:")
	fmt.Println(projects)
	distrib := DistrBudget(totalBudget, projects)
	fmt.Println("\nРаспределение бюджета между проектами:")
	for name, budget := range distrib {
		fmt.Printf("%s: %d единиц\n", name, budget)
	}
}

func DistrBudget(totalBudget int, projects []Project) map[string]int {
	var total int
	distrib := make(map[string]int)

	// общий приоритет всех проектов
	for _, project := range projects {
		total += project.Priority
	}

	// бюджет между проектами
	for _, project := range projects {
		// доля бюджета для каждого проекта
		budgetForProject := (project.Priority * totalBudget) / total
		distrib[project.Name] = budgetForProject
	}

	return distrib
}

type Project struct {
	Name     string
	Priority int
}

// У вас есть компания, работающая над несколькими проектами одновременно. Каждый проект имеет свою важность, выраженную в виде целочисленного приоритета. Ваша задача — написать программу на Go, которая будет распределять доступный общий бюджет между проектами с учетом их приоритета.

// Входные данные:
// Общий бюджет компании на проекты.
// Список проектов, где каждый проект имеет название и приоритет (целое число).
// Условия:
// Бюджет каждого проекта должен быть пропорционален его приоритету: чем выше приоритет, тем больше бюджета проект получает.
// Сумма бюджетов всех проектов не должна превышать общий бюджет.
// Результат должен быть выведен в виде списка названий проектов и соответствующих им сумм бюджета.
// Пример:
// Общий бюджет: 1000 единиц.

// Проекты:

// Проект A с приоритетом 5
// Проект B с приоритетом 3
// Проект C с приоритетом 2
// Решение:
// Сначала рассчитываем общую сумму приоритетов всех проектов.
// Для каждого проекта вычисляем долю от общего бюджета исходя из его приоритета.
// Пример:
// Общий бюджет: 1000 единиц.

// Проекты:

// Проект A с приоритетом 5
// Проект B с приоритетом 3
// Проект C с приоритетом 2
// Ожидаемый вывод:
// Проект A: 500 единиц
// Проект B: 300 единиц
// Проект C: 200 единиц
// type Project struct {
//     Name      string
//     Priority int
// }

// func distributeBudget(totalBudget int, projects []Project) map[string]int {
//     // Ваш код здесь
// }
