package person

import (
	"fmt"
)

type PersonInterface interface {
	GetInfo()
	SetAge(age int) error
	AddFriend(name string, age int) error
	ListFriends()
}
type Person struct {
	Name    string
	Age     int
	Friends map[string]int
}

func (p *Person) GetInfo() {
	fmt.Printf("Имя:%s, возраст: %d\n", p.Name, p.Age)
}

func (p *Person) SetAge(age int) error {
	if age < 0 {
		return fmt.Errorf("возраст не может быть отрицательный")
	}

	p.Age = age

	return nil
}

func (p *Person) AddFriend(name string, age int) error {
	if age < 0 || name == "" {
		return fmt.Errorf("введены некоректные данные")
	}
	if p.Friends == nil {
		p.Friends = make(map[string]int)
	}

	p.Friends[name] = age

	return nil
}

func (p *Person) ListFriends() {
	for name, age := range p.Friends {
		fmt.Printf("Имя друга:%s, возраст: %d\n", name, age)
	}
}
