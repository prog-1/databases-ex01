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
	fmt.Println(studentCountPerClass(db))
	fmt.Println(studentCountPerYear(db))
	fmt.Println(lessonsPerYear(db, 10))
	fmt.Println(examsPerClass(db, 10, "b"))
	fmt.Println(averageGradeForStudents(db, "Andrejs"))

}

func studentCountPerClass(db tables) map[string]int {
	m := make(map[string]int)
	for _, v := range db.groups.Data {
		class := db.classes.MustFindById(v.ClassID)
		m[strconv.Itoa(class.Year)+class.Mod]++
	}
	return m
}
func studentCountPerYear(db tables) map[int]int {
	m := make(map[int]int)
	for _, v := range db.groups.Data {
		class := db.classes.MustFindById(v.ClassID)
		m[class.Year]++
	}
	return m
}
func lessonsPerYear(db tables, year int) []string {
	var m []string
	for _, v := range db.timetable.Data {
		class := db.classes.MustFindById(v.ClassID)
		if class.Year == year {
			lesson := db.lessons.FindById(v.LessonID)
			m = append(m, lesson.Name)
		}
	}
	return m
}
func examsPerClass(db tables, year int, mod string) []string {
	var m []string
	for _, h := range db.exams.Data {
		for _, v := range db.timetable.Data {
			if h.LessonID == v.LessonID {
				class := db.classes.MustFindById(v.ClassID)
				if class.Year == year && class.Mod == mod {
					lesson := db.lessons.FindById(v.LessonID)
					m = append(m, lesson.Name)
				}
			}
		}
	}
	return m
}
func averageGradeForStudents(db tables, name string) float64 {
	var cnt, sum int
	for _, v := range db.exams.Data {
		v1 := db.students.FindStudentById(v.StudentID)
		if v1.Name == name {
			sum += v.Grade
		}

	}
	return float64(sum) / float64(cnt)
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
