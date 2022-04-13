package database

type Exam struct {
	ID        int
	StudentID int
	LessonID  int
	Grade     int
}

type ExamTable struct {
	Data []Exam
}

func (db *ExamTable) Insert(sID, lID, grade int) Exam {
	e := Exam{
		ID:        len(db.Data) + 1,
		StudentID: sID,
		LessonID:  lID,
		Grade:     grade,
	}
	db.Data = append(db.Data, e)
	return e
}
