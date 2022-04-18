package database

// Lesson represents a row storing information about a school lesson.
type Exams struct {
	ID        int
	StudentID int
	LessonID  int
	Grade     int
}

// LessonTable is a class database.
type ExamsTable struct {
	Data []Exams
}

func (db *ExamsTable) ExamFindById(id int) Exams {
	for _, c := range db.Data {
		if c.ID == id {
			return c
		}
	}
	panic("not found")
}

func (db *ExamsTable) Insert(lID, sID, grade int) Exams {
	e := Exams{
		LessonID:  lID,
		StudentID: sID,
		Grade:     grade,
	}
	db.Data = append(db.Data, e)
	return e
}

func (db *ExamsTable) AddExam(sID int, lID int, grade ...int) {
	for _, l := range grade {
		db.Insert(sID, lID, l)
	}
}
