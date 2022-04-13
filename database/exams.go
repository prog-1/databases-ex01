package database

type Exams struct {
	ID        int
	StudentID int
	LessonID  int
	Grade     int
}

// GroupTable is a group database.
type ExamTable struct {
	Data []Exams
}

func (db *ExamTable) Insert(sID int, lID int, grade int) Exams {
	tt := Exams{
		ID:        len(db.Data) + 1,
		StudentID: sID,
		Grade:     grade,
		LessonID:  lID,
	}
	db.Data = append(db.Data, tt)
	return tt
}

// AddTimetable adds rows for several lessons to the timetable.
func (db *ExamTable) AddExam(sID int, lID int, grade ...int) {
	for _, le := range grade {
		db.Insert(sID, lID, le)
	}
}
