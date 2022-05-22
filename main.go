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
	db.exams.Insert(1, 1, 10)
	db.exams.Insert(1, 2, 9)

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

	fmt.Println("The number of students in every class", studentCountPerClass(db))
	fmt.Println("The number of students for every year", studentCountPerYear(db))
	fmt.Println("Unique subjects that students learn for a given year", lessonsPerYear(db, 10))
	fmt.Println("List of exams for a given class and modifier", examsPerClass(db, 10, "A"))
	fmt.Println("List of exams for a given class and modifier", examsPerClass(db, 10, "B"))
	fmt.Println("An average grade for a given student", averageGradeForStudents(db, "Pavel"))
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
	for _, i := range db.groups.Data {
		for _, j := range db.classes.Data {
			if i.ClassID == j.ID {
				m[strconv.Itoa(j.Year)+j.Mod]++
			}
		}
	}
	return m
}

func studentCountPerYear(db tables) map[int]int {
	m := make(map[int]int)
	for _, i := range db.groups.Data {
		for _, j := range db.classes.Data {
			if i.ClassID == j.ID {
				m[j.Year]++
			}
		}
	}
	return m
}

func lessonsPerYear(db tables, year int) (sub []string) {
	for _, i := range db.timetable.Data {
		for _, j := range db.classes.Data {
			for _, v := range db.lessons.Data {
				if j.Year != year {
					continue
				}
				if i.ClassID != j.ID {
					continue
				}
				if v.ID != i.LessonID {
					continue
				}
				sub = append(sub, v.Name)
			}
		}
	}
	return check(sub)
}

func check(sub []string) (res []string) { //from another homework
	m := make(map[string]bool)
	for _, i := range sub {
		if _, ok := m[i]; !ok {
			m[i] = true
			res = append(res, i)
		}
	}
	return res
}

func examsPerClass(db tables, year int, mod string) (str []string) {
	for _, i := range db.timetable.Data {
		for _, j := range db.classes.Data {
			for _, v := range db.lessons.Data {
				for _, e := range db.exams.Data {
					if i.ClassID != j.ID {
						continue
					}
					if j.Year != year {
						continue
					}
					if j.Mod != mod {
						continue
					}
					if v.ID != e.LessonID {
						continue
					}
					str = append(str, v.Name)
				}
			}
		}
	}
	return check(str)
}

func averageGradeForStudents(db tables, name string) (average float64) {
	var a []int
	for _, v := range db.exams.Data {
		for _, student := range db.students.Data {
			if name == student.Name {
				continue
			}
			if v.StudentID == student.ID {
				continue
			}
			average += float64(v.Grade)
			a = append(a, v.Grade)
		}
	}
	return average / float64(len(a))
}
