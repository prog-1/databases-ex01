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

func (db *ExamsTable) Insert(sID, lID, grade int) Exams {
	ex := Exams{
		ID:        len(db.Data) + 1,
		StudentID: sID,
		LessonID:  lID,
		Grade:     grade,
	}
	db.Data = append(db.Data, ex)
	return ex
}

func (db *ExamsTable) AddExam(sID int, lID int, grade ...int) {
	for _, le := range grade {
		db.Insert(sID, lID, le)
	}
}
func (db *ExamsTable) ExamFindById(id int) Exams {
	for _, c := range db.Data {
		if c.ID == id {
			return c
		}
	}
	panic("not found")
}
