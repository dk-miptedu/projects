package cmd

import (
	"bufio"
	"dz-04_Topic05_BookLibrary/models"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("Library Manage App")
	library := models.Library{}

	// Добавление книг
	//library.AddBook(models.Book{Title: "Война и мир", Author: "Лев Толстой", Publication: 1869})
	//library.AddBook(models.Book{Title: "1984", Author: "Джордж Оруэлл", Publication: 1949})

	// Отображение списка книг
	//library.DisplayBooks()

	// Чтение книги
	//ReadBook(library.Books[0])

	scanner := bufio.NewScanner(os.Stdin)
	var publication int
	var err error
	for {
		fmt.Println("\nВведите команду (add, remove, display, read, exit):")
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "add":
			fmt.Println("Введите Название, Автора и Год издания через запятую:")
			scanner.Scan()
			bookInfo := strings.Split(scanner.Text(), ",")
			if len(bookInfo) != 3 {
				fmt.Println("Неверный формат ввода. Попробуйте снова.")
				continue
			} else {
				// Преобразование строки года в целое число
				publication, err = strconv.Atoi(bookInfo[2])

				if err != nil {
					fmt.Println("Ошибка при вводе Год издания")
					continue
				}

			}

			library.AddBook(models.Book{Title: bookInfo[0], Author: bookInfo[1], Publication: publication})
		case "remove":
			fmt.Println("Введите название книги для удаления:")
			scanner.Scan()
			library.RemoveBook(scanner.Text())
		case "display":
			library.DisplayBooks()
		case "read":
			fmt.Println("Введите название книги для чтения:")
			scanner.Scan()
			title := scanner.Text()
			for _, book := range library.Books {
				if book.Title == title {
					models.ReadBook(book)
					break
				}
			}
		case "exit":
			return
		default:
			fmt.Println("Неизвестная команда.")
		}
	}

}
