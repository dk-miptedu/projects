package main

import (
	"fmt"

	"web20240227/person"
)

func main() {
	person := person.Person{
		Name: "VAsya",
		Age:  30,
	}
	person.GetInfo()
	err := person.SetAge(35)
	if err != nil {
		fmt.Println(err)
	}

	err = person.AddFriend("Pet", 25)

	if err := person.AddFriend("petty", 26); err != nil {
		fmt.Println(err)
	}
	if err := person.AddFriend("Drytty", 36); err != nil {
		fmt.Println(err)
	}
	person.ListFriends()
}
