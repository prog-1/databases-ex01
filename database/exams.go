package database

type Exams struct {
	ID        int
	StudentID int
	LessonID  int
	Grade     int
}

type ExamTable struct {
	Data []Exams
}

func (db *ExamTable) Insert(stID int, lesID int, grade int) Exams {
	e := Exams{
		ID:        len(db.Data) + 1,
		StudentID: stID,
		LessonID:  lesID,
		Grade:     grade,
	}
	db.Data = append(db.Data, e)
	return e
}

func (db *ExamTable) AddExam(stID int, lesID int, grade ...int) {
	for _, v := range grade {
		db.Insert(stID, lesID, v)
	}
}
