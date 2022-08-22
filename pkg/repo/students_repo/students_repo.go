package students_repo

import (
	"fmt"
	"module/pkg/student"
)

type StudentsStorage struct {
	database map[string]*student.Student
}

func NewStudentsStorage() *StudentsStorage {
	return &StudentsStorage{
		database: make(map[string]*student.Student),
	}
}

func (ss StudentsStorage) Put(student *student.Student) error {
	if student == nil {
		return fmt.Errorf("No such a student: struct is empty")
	}

	if student.Name == "" {
		return fmt.Errorf("Can't include student: name is empty")
	}

	ss.database[student.Name] = student
	return nil
}

func (ss StudentsStorage) Get() []*student.Student {
	res := []*student.Student{}
	for _, v := range ss.database {
		res = append(res, v)
	}

	return res
}
