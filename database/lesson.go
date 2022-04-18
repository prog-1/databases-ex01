package database

// Lesson represents a row storing information about a school lesson.
type Lesson struct {
	ID   int
	Name string
}

// LessonTable is a class database.
type LessonTable struct {
	Data []Lesson
}

func (db *LessonTable) FindById(id int) Lesson {
	for _, c := range db.Data {
		if c.ID == id {
			return c
		}
	}
	panic("not found")
}

// Insert adds a new row to the lesson database.
func (db *LessonTable) Insert(name string) Lesson {
	l := Lesson{
		ID:   len(db.Data) + 1,
		Name: name,
	}
	db.Data = append(db.Data, l)
	return l
}
