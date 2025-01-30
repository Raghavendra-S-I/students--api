package storage

import "github.com/Raghavendra/students-api/internal/types"

type Storage interface {

	// Create a new student
	CreateStudent(name string, email string, age int) (int64, error)

	GetStudents(id int64) (types.Student, error)
}
