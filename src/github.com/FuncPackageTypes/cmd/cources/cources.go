package cources

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var courseStudents map[string]int

func addStudent(course string) {
	if courseStudents == nil {
		courseStudents = make(map[string]int)
	}
	caseTitle := cases.Title(language.English)
	course = caseTitle.String(strings.ToLower(course))
	courseStudents[course]++
}
func getStudentsCount(course string) int {
	return courseStudents[course]
}
func Booking() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Введите команду (add, count, exit):")
		scanner.Scan()
		command := strings.TrimSpace(scanner.Text())

		switch command {
		case "add":
			fmt.Println("Введите название курса:")
			scanner.Scan()
			course := strings.TrimSpace(scanner.Text())
			addStudent(course)
			fmt.Println("Студент добавлен на курс", course)
		case "count":
			for i, v := range courseStudents {
				fmt.Printf("Количество студентов на курсе '%s': %d\n", i, v)
			}
		case "exit":
			return
		default:
			fmt.Println("Неизвестная команда. Попробуйте снова.")
		}
	}
}
