package models

import (
	"dz-04_Topic05_BookLibrary/interfaces"
	"fmt"
)

type Book struct {
	Title       string
	Author      string
	Publication int
}

type Library struct {
	Books []Book
}

// Реализация метода AddBook
func (l *Library) AddBook(book Book) {
	l.Books = append(l.Books, book)
}

// Реализация метода RemoveBook
func (l *Library) RemoveBook(title string) {
	for i, book := range l.Books {
		if book.Title == title {
			l.Books = append(l.Books[:i], l.Books[i+1:]...)
			fmt.Printf("Книга %s удалена из списка\n", book.Title)
			l.DisplayBooks()
			break
		}
	}
}

// Реализация метода DisplayBooks
func (l *Library) DisplayBooks() {
	fmt.Println("Список книг в библиотеке:")
	for _, book := range l.Books {
		fmt.Printf("Название: %s, Автор: %s, Год издания: %d\n", book.Title, book.Author, book.Publication)
	}
}

func ReadBook(r interfaces.Readable) {
	r.Read()
}

// Реализация метода Read() для структуры Book
func (b Book) Read() {
	fmt.Printf("Читатель начал читать книгу \"%s\" автора %s\n", b.Title, b.Author)
}
