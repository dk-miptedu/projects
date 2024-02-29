package dictionaries

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var translations map[string]string

func addTranslation(word, translation string) {
	if translations == nil {
		translations = make(map[string]string)
	}
	translations[word] = translation
}

func findTranslation(word string) string {
	if translation, exists := translations[word]; exists {
		return translation
	}
	return "Перевод не найден"
}

func DictManage() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Введите команду (add, find, exit):")
		scanner.Scan()
		command := scanner.Text()

		switch command {
		case "add":
			fmt.Println("Введите слово и его перевод через пробел:")
			scanner.Scan()
			input := scanner.Text()
			parts := strings.SplitN(input, " ", 2)
			if len(parts) != 2 {
				fmt.Println("Некорректный ввод. Попробуйте снова.")
				continue
			}
			addTranslation(parts[0], parts[1])
			fmt.Println("Перевод добавлен.")
		case "find":
			fmt.Println("Введите слово для поиска его перевода:")
			scanner.Scan()
			word := scanner.Text()
			translation := findTranslation(word)
			fmt.Printf("Перевод: %s\n", translation)
		case "exit":
			return
		default:
			fmt.Println("Неизвестная команда. Попробуйте снова.")
		}
	}
}
