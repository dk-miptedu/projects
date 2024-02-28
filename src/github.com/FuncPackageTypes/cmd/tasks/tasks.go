package tasks

var TaskHomeWork = [4]string{
	"написать программу, которая принимает строку от пользователя и создаёт слайс, содержащий каждый символ этой строки.",
	"создать структуру для представления информации о студентах (имя, возраст, оценки и т. д.). Реализовать функцию, которая добавляет студентов в слайс и выводит список студентов.",
	"создать map для хранения перевода слов с одного языка на другой. У пользователя должна быть возможность добавлять и искать переводы.",
	"используйте map, чтобы учитывать количество студентов по каждому курсу в университете.",
}

func GetTaskHomeWork() map[int]string {
	TaskHomeWorkMap := make(map[int]string)
	for i, task := range TaskHomeWork {
		TaskHomeWorkMap[i+1] = task
	}
	return TaskHomeWorkMap
}