package database

// Student represents a row storing information about a student.
type Student struct {
	ID   int
	Name string
}

// StudentTable is a class database.
type StudentTable struct {
	Data []Student
}

// Insert adds a new row to the student database.
func (db *StudentTable) Insert(name string) Student {
	s := Student{
		ID:   len(db.Data) + 1,
		Name: name,
	}
	db.Data = append(db.Data, s)
	return s
}
func (db *StudentTable) FindStudentById(id int) Student {
	for _, c := range db.Data {
		if c.ID == id {
			return c
		}
	}
	panic("not found")
}
