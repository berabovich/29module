package main

import (
	"bufio"
	"fmt"
	"os"
	"test/pkg/storage"
	"test/pkg/student"
)

func main() {

	m := make(map[string]*student.Student)
	studentStorage := storage.Storage{}
	fmt.Println("Введите имя, возраст и курс студента!")
	for {
		scanner := bufio.NewScanner(os.Stdin)
		if !scanner.Scan() {
			break
		}
		newStudent := scanner.Text()
		studentStorage.Put(m, newStudent)
	}
	studentStorage.Get(m)
}
