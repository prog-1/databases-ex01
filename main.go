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
	db.exams.Insert(1, 2, 9)
	db.exams.Insert(2, 3, 3)
	db.exams.Insert(2, 1, 9)
	db.exams.Insert(1, 1, 10)
	db.exams.Insert(5, 2, 5)
	db.exams.Insert(6, 2, 7)

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
	fmt.Println(studentCountPerClass(db))
	fmt.Println(studentCountPerYear(db))
	fmt.Println(lessonsPerYear(db, 10))
	fmt.Println(examsPerClass(db, 10, "A"))
	fmt.Println(averageGradeForStudents(db, "Pavels"))

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
	scpc := make(map[string]int)
	for _, gr := range db.groups.Data {
		for _, cl := range db.classes.Data {
			if gr.ClassID == cl.ID {
				in := fmt.Sprint(cl.Year) + fmt.Sprint(cl.Mod)
				scpc[in]++
			}
		}
	}
	return scpc
}

func studentCountPerYear(db tables) map[int]int {
	scpy := make(map[int]int)
	for _, gr := range db.groups.Data {
		for _, cl := range db.classes.Data {
			if gr.ClassID == cl.ID {
				scpy[cl.Year]++
			}
		}
	}
	return scpy
}

func lessonsPerYear(db tables, year int) (lpy []string) {
	var i []int
	var duplicate bool
	for _, tt := range db.timetable.Data {
		for _, cl := range db.classes.Data {
			if tt.ClassID == cl.ID && cl.Year == year {
				for _, ls := range db.lessons.Data {
					if tt.LessonID == ls.ID {
						if len(i) != 0 {
							for _, v := range i {
								if tt.LessonID == v {
									duplicate = true
								}
							}
							if !duplicate {
								lpy = append(lpy, ls.Name)
								i = append(i, tt.LessonID)
							}
						} else {
							lpy = append(lpy, ls.Name)
							i = append(i, tt.LessonID)
						}
					}
				}

			}
		}
	}
	return lpy
}

func examsPerClass(db tables, year int, mod string) (epc []string) {
	var duplicate bool
	for _, ex := range db.exams.Data {
		for _, st := range db.students.Data {
			if ex.StudentID == st.ID {
				for _, gr := range db.groups.Data {
					if st.ID == gr.StudentID {
						for _, cl := range db.classes.Data {
							if gr.ClassID == cl.ID {
								if cl.Mod == mod && cl.Year == year {
									for _, ls := range db.lessons.Data {
										if ls.ID == ex.LessonID {
											for _, v := range epc {
												if v == ls.Name {
													duplicate = true
												}
											}
											if !duplicate {
												epc = append(epc, ls.Name)
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return epc
}

func averageGradeForStudents(db tables, name string) (avg float64) {
	var grades []int
	var a float64
	for _, ex := range db.exams.Data {
		for _, st := range db.students.Data {
			if ex.StudentID == st.ID && st.Name == name {
				grades = append(grades, ex.Grade)
			}
		}
	}
	for _, v := range grades {
		a = a + float64(v)
	}
	avg = a / float64(len(grades))
	return avg
}
