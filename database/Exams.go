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

func (db *ExamsTable) ExamFindById(id int) Exams {
	for _, c := range db.Data {
		if c.ID == id {
			return c
		}
	}
	panic("not found")
}
