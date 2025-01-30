package storage

type Storage interface {

	// Create a new student
	CreateStudent(name string, email string, age int) (int64, error)
}
