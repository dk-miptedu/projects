package main

import (
	"FuncPackageTypes/cmd/cources"
	"FuncPackageTypes/cmd/dictionaries"
	"FuncPackageTypes/cmd/slices"
	"FuncPackageTypes/cmd/students"
	"FuncPackageTypes/cmd/tasks"
	"fmt"
)

func main() {
	homeTasks := tasks.GetTaskHomeWork()
	for i, homeTask := range homeTasks {
		fmt.Println("=====================")
		fmt.Printf("Задание %d: %s\n", i, homeTask)
		switch i {
		case 1:
			slices.SplitInputStrings()
		case 2:
			students.StudentsTest()
		case 3:
			dictionaries.DictManage()
		case 4:
			cources.Booking()
			//default:
			//	fmt.Println("Неизвестная задача, Выход")
			//	break
		}

	}
}
