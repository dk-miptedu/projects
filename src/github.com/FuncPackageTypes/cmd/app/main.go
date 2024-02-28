package main

import (
	"FuncPackageTypes/cmd/person"
	"FuncPackageTypes/cmd/slices"
	"FuncPackageTypes/cmd/students"
	"FuncPackageTypes/cmd/tasks"
	"fmt"
)

func main() {
	homeTasks := tasks.GetTaskHomeWork()
	for i := 0; i <= len(homeTasks); i++ {

		fmt.Println(homeTasks[i])

	}

	fmt.Println("**Задание 1**")
	slices.SplitInputStrings()

	fmt.Println("Задание 2")
	students.StudentsTest()

	fmt.Println("Задание 3")

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
