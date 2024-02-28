package students

import "fmt"

type Student struct {
	Name   string
	Age    int
	Grades []float64
}

type StudentList struct {
	ListItem []Student
}

func (f *StudentList) AddStudent(newStudent Student) {
	f.ListItem = append(f.ListItem, newStudent)
}

func (f *StudentList) PrintStudents() {
	for _, student := range f.ListItem {
		fmt.Printf("Имя: %s, Возраст: %d, Оценки: %v\n", student.Name, student.Age, student.Grades)
	}
}

func StudentsTest() {
	var userAddStudents StudentList
	userAddStudents.AddStudent(Student{Name: "Алексей", Age: 20, Grades: []float64{4.5, 4.0, 5.0}})
	userAddStudents.AddStudent(Student{Name: "Мария", Age: 19, Grades: []float64{5.0, 5.0, 5.0}})

	userAddStudents.PrintStudents()
}
