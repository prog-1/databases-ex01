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

func (db *ExamsTable) Insert(cID int, grade int, lID int) Exams {
	tt := Exams{
		ID:        len(db.Data) + 1,
		StudentID: cID,
		Grade:     grade,
		LessonID:  lID,
	}
	db.Data = append(db.Data, tt)
	return tt
}

func (db *ExamsTable) AddTimetable(cID int, lID int, grade ...int) {
	for _, le := range grade {
		db.Insert(cID, lID, le)
	}
}
