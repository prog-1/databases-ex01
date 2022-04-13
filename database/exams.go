package database

type Exams struct {
	ID        int
	StudentID int
	LessonID  int
	Grade     int
}

type ExamsTable struct {
	Data []Exams
}

func (db *ExamsTable) Insert(sID, lID, g int) Exams {
	e := Exams{
		ID:        len(db.Data) + 1,
		StudentID: sID,
		LessonID:  lID,
		Grade:     g,
	}
	db.Data = append(db.Data, e)
	return e
}

func (db *ExamsTable) AddExam(sID int, lID int, grade ...int) {
	for _, le := range grade {
		db.Insert(sID, lID, le)
	}
}
