package cmd

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// NewRand создаёт и возвращает новый экземпляр rand.Rand, инициализированный текущим временем
func NewRand() *rand.Rand {
	source := rand.NewSource(time.Now().UnixNano())
	return rand.New(source)
}

func randomString(n int, r *rand.Rand) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[r.Intn(len(letters))]
	}
	return string(s)
}

func GenUserData(r *rand.Rand) (string,string) {
	user := randomString(4, r) // Генерируем случайное имя пользователя длиной 10 символов
	domain := randomString(4, r) // Генерируем случайное доменное имя длиной 5 символов
	email := fmt.Sprintf("%s@%s.org", user, domain)
	return user, email
}

func main() {
	r := NewRand() // Создаём новый экземпляр rand.Rand
	email, user := generateRandomEmail(r)
	fmt.Println("Email:", email)
	fmt.Println("User:", user)

}
