package slices

import (
	"bufio"
	"fmt"
	"os"
)

func SplitInputStrings() {
	var userInput string
	var userInputchars []string
	fmt.Println("Введите строку и нажмите Enter:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userInput = scanner.Text()
	for i := 0; i < len(userInput); i++ {
		userInputchars = append(userInputchars, string(userInput[i]))
	}
	fmt.Println("Слайс, содержащий каждый символ строки:", userInputchars)
}
