package main

import (
	"databases-ex01/database"
	"fmt"
	"time"
)

type tables struct {
	students  database.StudentTable
	classes   database.ClassTable
	groups    database.GroupTable
	lessons   database.LessonTable
	timetable database.TimetableTable
	exams     database.ExamsTable
}

func main() {
	var db tables

	db.groups.AddStudentsToClass(
		db.classes.Insert(10, "B"),
		db.students.Insert("Jaroslavs"),
		db.students.Insert("Pavels"),
	)

	db.groups.AddStudentsToClass(
		db.classes.Insert(10, "A"),
		db.students.Insert("Alina"),
		db.students.Insert("Andrejs"),
		db.students.Insert("Antons"),
		db.students.Insert("Valerija"),
	)

	maths := db.lessons.Insert("Mathematics")
	programming := db.lessons.Insert("Programming")
	sports := db.lessons.Insert("Sports")

	db.timetable.AddTimetable(
		db.classes.MustFind(10, "A"),
		time.Monday,
		maths,
		programming,
	)
	db.timetable.AddTimetable(
		db.classes.MustFind(10, "A"),
		time.Tuesday,
		sports,
	)

	db.timetable.AddTimetable(
		db.classes.MustFind(10, "B"),
		time.Tuesday,
		maths,
	)

	fmt.Println("We have", len(db.students.Data), "unique students in our school")
	fmt.Println("We have", len(db.classes.Data), "unique classes in our school")

	fmt.Println("On Tuesday Jaroslavs has:")
	for _, l := range studentDayTimetable(db, "Jaroslavs", time.Tuesday) {
		fmt.Println(" ", l.Name)
	}
	fmt.Println("On Monday Pavels has:")
	for _, l := range studentDayTimetable(db, "Pavels", time.Monday) {
		fmt.Println(" ", l.Name)
	}

	fmt.Println("Pavels has:")
	for d, lessons := range studentWeekTimetable(db, "Pavels") {
		fmt.Println(" *", d)
		for _, l := range lessons {
			fmt.Println("   - ", l.Name)
		}
	}

	fmt.Println("Alina has:")
	for d, lessons := range studentWeekTimetable(db, "Alina") {
		fmt.Println(" *", d)
		for _, l := range lessons {
			fmt.Println("   - ", l.Name)
		}
	}

	fmt.Println("Student  per class", studentCountPerClass(db))
	fmt.Println("Student  per year", studentCountPerYear(db))
	fmt.Println("Lessons per year", lessonsPerYear(db, 10))
	fmt.Println("Exams per class", examsPerClass(db, 10, "B"))
	fmt.Println("Average grade for the student", averageGradeForStudents(db, "Pavels"))

}

func studentDayTimetable(db tables, name string, day time.Weekday) (res []database.Lesson) {
	for _, s := range db.students.Data {
		if s.Name != name {
			continue
		}

		for _, g := range db.groups.Data {
			if g.StudentID != s.ID {
				continue
			}

			for _, tt := range db.timetable.Data {
				if tt.ClassID != g.ClassID {
					continue
				}
				if tt.Day != day {
					continue
				}

				for _, l := range db.lessons.Data {
					if l.ID != tt.LessonID {
						continue
					}

					res = append(res, l)
				}
			}
		}
	}
	return
}

func studentWeekTimetable(db tables, name string) map[time.Weekday][]database.Lesson {
	res := make(map[time.Weekday][]database.Lesson)
	for _, s := range db.students.Data {
		if s.Name != name {
			continue
		}

		for _, g := range db.groups.Data {
			if g.StudentID != s.ID {
				continue
			}

			for _, tt := range db.timetable.Data {
				if tt.ClassID != g.ClassID {
					continue
				}

				for _, l := range db.lessons.Data {
					if l.ID != tt.LessonID {
						continue
					}

					res[tt.Day] = append(res[tt.Day], l)
				}
			}
		}
	}
	return res
}

func studentCountPerClass(db tables) map[string]int {
	m := make(map[string]int)
	var st string
	for _, i := range db.classes.Data {
		for _, e := range db.groups.Data {
			if i.ID == e.ClassID {
				st = fmt.Sprint(i.Year) + i.Mod
				m[st]++
			}
		}
	}
	return m
}

func studentCountPerYear(db tables) map[int]int {
	res := make(map[int]int)
	for _, i := range db.groups.Data {
		for _, j := range db.classes.Data {
			if i.ClassID == j.ID {
				res[j.Year]++
			}
		}
	}
	return res
}

func lessonsPerYear(db tables, year int) (str []string) {
	for _, tt := range db.timetable.Data {
		for _, cl := range db.classes.Data {
			for _, les := range db.lessons.Data {
					for _, v := range en {
						if v == ls.Name {
							duplicate = true
						}
					}
					if !duplicate {
						en = append(en, ls.Name)
					}
				}
			}
		}
	}

func examsPerClass(db tables, year int, mod string) (en []string) {
	for _, tt := range db.timetable.Data {
		for _, cl := range db.classes.Data {
			for _, les := range db.lessons.Data {
				for _, ex := range db.exams.Data {
					for _, v := range en {
						if v == ls.Name {
							duplicate = true
						}
					}
					if !duplicate {
						en = append(en, ls.Name)
					}
					}
				}
			}
		}
	}
	return en
}

func averageGradeForStudents(db tables, name string) (average float64) {
	var g []int
	for _, exam := range db.exams.Data {
		for _, student := range db.students.Data {
				if exam.StudentID == st.ID && st.Name == name {
					g = append(g, exam.Grade)
				}
		}
		return average / float64(len(g))
	}
}

