package storage

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"test/pkg/student"
)

type Storage struct {
	db map[string]*student.Student
}

func (db Storage) Put(students map[string]*student.Student,newStudent string) {
	s := strings.Fields(newStudent)

	newName, age, grade := s[0], s[1], s[2]

	newAge, err := strconv.Atoi(age)
	if err != nil {
		err = errors.New("Неверный ввод возраста")
		fmt.Println(err)
		return
	}
	newGrade, err := strconv.Atoi(grade)
	if err != nil {
		err = errors.New("Неверный ввод курса")
		fmt.Println(err)
		return
	}
	students[newStudent] = &student.Student {
		Name:  newName,
		Age:   newAge,
		Grade: newGrade,
	}
}

func (db Storage) Get(m map[string]*student.Student) {
	fmt.Println("Студенты из хранилища:")
	for key := range m {
		fmt.Println(key)
	}
}
