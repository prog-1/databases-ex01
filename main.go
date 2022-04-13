package main

import (
	"databases-ex01/database"
	"fmt"
	"strconv"
	"time"
)

type tables struct {
	students  database.StudentTable
	classes   database.ClassTable
	groups    database.GroupTable
	lessons   database.LessonTable
	timetable database.TimetableTable
	exams     database.ExamTable
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
	fmt.Println()

}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
func examsPerClass(db tables, year int, mod string) []string {
	dictionary := make(map[int]string)
	for _, v := range db.lessons.Data {
		dictionary[v.ID] = v.Name
	}
	var exams []string
	for _, v := range db.classes.Data {
		if v.Mod == mod && v.Year == year {
			for _, v1 := range db.groups.Data {
				if v1.ClassID == v.ID {
					for _, v2 := range db.students.Data {
						if v2.ID == v1.StudentID {
							for _, v3 := range db.exams.Data {
								if v2.ID == v3.StudentID && contains(exams, dictionary[v3.LessonID]) {
									exams = append(exams, dictionary[v3.LessonID])
								}
							}
						}
					}
				}
			}
		}
	}
	return exams
}
func averageGradeForStudents(db tables, name string) float64 {
	for _, v := range db.students.Data {
		if v.Name == name {
			var avg float64
			var i float64
			for _, v2 := range db.exams.Data {
				if v.ID == v2.StudentID {
					avg += float64(v2.Grade)
					i++
				}
			}
			return avg / i

		}
	}
	return 0
}
func lessonsPerYear(db tables, year int) []string {
	var lessons []string
	for _, v := range db.classes.Data {
		if v.Year != year {
			continue
		}
		for _, v1 := range db.timetable.Data {
			if v.ID == v1.ClassID {
				for _, v2 := range db.lessons.Data {
					if v2.ID == v1.LessonID && !contains(lessons, v2.Name) {
						lessons = append(lessons, v2.Name)
					}
				}

			}
		}
	}
	return lessons
}

func studentCountPerYear(db tables) map[int]int {
	dictionary := make(map[int]int)
	for _, v := range db.classes.Data {
		dictionary[v.ID] = v.Year
	}
	data := make(map[int]int)
	for _, v := range db.groups.Data {
		data[dictionary[v.ClassID]]++
	}
	return data
}

func studentCountPerClass(db tables) map[string]int {
	dictionary := make(map[int]string)
	for _, v := range db.classes.Data {
		dictionary[v.ID] = strconv.Itoa(v.Year) + v.Mod
	}
	data := make(map[string]int)
	for _, v := range db.groups.Data {
		data[dictionary[v.ClassID]]++
	}
	return data
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
