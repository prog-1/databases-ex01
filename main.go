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

	db.exams.Insert(1, 1, 10)
	db.exams.Insert(1, 2, 10)
	db.exams.Insert(2, 1, 10)
	db.exams.Insert(2, 2, 10)
	db.exams.Insert(4, 1, 7)
	db.exams.Insert(4, 2, 8)
	db.exams.Insert(5, 1, 6)
	db.exams.Insert(5, 2, 6)
	db.exams.Insert(5, 3, 9)

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

	fmt.Println("Student count per class:")
	for c, count := range studentCountPerClass(db) {
		fmt.Printf(" %s: %v\n", c, count)
	}

	fmt.Println("Student count per year:")
	for y, count := range studentCountPerYear(db) {
		fmt.Printf(" %v: %v\n", y, count)
	}

	fmt.Println("Lessons per year 10:")
	for _, v := range lessonsPerYear(db, 10) {
		fmt.Printf(" %s\n", v)
	}

	fmt.Println("Exams per class 10B:")
	for _, v := range examsPerClass(db, 10, "B") {
		fmt.Printf(" %s\n", v)
	}

	fmt.Println("Average grade for student Andrejs: ", averageGradeForStudents(db, "Andrejs"))
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
	res := make(map[string]int)
	for _, g := range db.groups.Data {
		for _, c := range db.classes.Data {
			if g.ClassID != c.ID {
				continue
			}

			res[fmt.Sprintf("%v%s", c.Year, c.Mod)]++
		}
	}
	return res
}

func studentCountPerYear(db tables) map[int]int {
	res := make(map[int]int)
	for _, g := range db.groups.Data {
		for _, c := range db.classes.Data {
			if g.ClassID != c.ID {
				continue
			}

			res[c.Year]++
		}
	}
	return res
}

func lessonsPerYear(db tables, year int) []string {
	var res []string
	for _, c := range db.classes.Data {
		if c.Year != year {
			continue
		}

		for _, tt := range db.timetable.Data {
			if tt.ClassID != c.ID {
				continue
			}

			for _, l := range db.lessons.Data {
				if l.ID != tt.LessonID {
					continue
				}

				res = append(res, l.Name)
			}
		}
	}
	return removeDuplicates(res)
}

func removeDuplicates(s []string) []string {
	var res []string
	keys := make(map[string]bool)
	for _, entry := range s {
		if _, v := keys[entry]; !v {
			keys[entry] = true
			res = append(res, entry)
		}
	}
	return res
}

func examsPerClass(db tables, year int, mod string) []string {
	var res []string
	for _, g := range db.groups.Data {
		if g.ClassID != db.classes.MustFind(year, mod).ID {
			continue
		}

		for _, e := range db.exams.Data {
			for _, s := range db.students.Data {
				if s.ID != e.StudentID {
					continue
				}
				if s.ID != g.StudentID {
					continue
				}

				for _, l := range db.lessons.Data {
					if l.ID != e.LessonID {
						continue
					}

					res = append(res, l.Name)
				}
			}
		}
	}
	return removeDuplicates(res)
}

func averageGradeForStudents(db tables, name string) float64 {
	var res, gradeCount float64
	for _, s := range db.students.Data {
		if s.Name != name {
			continue
		}

		for _, e := range db.exams.Data {
			if e.StudentID != s.ID {
				continue
			}

			res += float64(e.Grade)
			gradeCount++
		}
	}
	return res / gradeCount
}
