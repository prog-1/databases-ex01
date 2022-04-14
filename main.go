package main

import (
	"databases-ex01/database"
	"fmt"
	"strings"
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

	studentid, lessonid := FindStudentAndLessonID(db, "Alina", "Mathematics")
	db.exams.Insert(studentid, lessonid, 9)
	studentid, lessonid = FindStudentAndLessonID(db, "Alina", "Programming")
	db.exams.Insert(studentid, lessonid, 8)

	// studentid, lessonid = FindStudentAndLessonID(db, "Vladimirs", "English") // to test 'panic()'
	// db.exams.Insert(studentid, lessonid, 10)

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

	studentcountperclass := studentCountPerClass(db)
	for class, studentcount := range studentcountperclass {
		fmt.Printf("Class %v has %v students\n", class, studentcount)
	}

	studentcountperyear := studentCountPerYear(db)
	for year, studentcount := range studentcountperyear {
		fmt.Printf("There are %v students for year %v\n", studentcount, year)
	}

	year := 10

	lessonsperyear := lessonsPerYear(db, year)
	fmt.Printf("There are following subjects that students learn for year %v:\n", year)
	for _, lesson := range lessonsperyear {
		fmt.Println(" *", lesson)
	}

	mod := "A"
	class := fmt.Sprint(year) + mod

	examsperclass := examsPerClass(db, year, mod)
	fmt.Printf("There are following exams in class %v:\n", class)
	for _, exam := range examsperclass {
		fmt.Println(" *", exam)
	}

	name := "Alina"

	averagegrade := averageGradeForStudents(db, name)
	fmt.Printf("%v's average grade is %v\n", name, averagegrade)

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
	var s string
	for _, i := range db.classes.Data {
		s = fmt.Sprint(i.Year) + i.Mod
		for _, e := range db.groups.Data {
			if i.ID == e.ClassID {
				m[s]++
			}
		}
	}
	return m
}

func studentCountPerYear(db tables) map[int]int {
	m := make(map[int]int)
	for _, i := range db.classes.Data {
		for _, e := range db.groups.Data {
			if i.ID == e.ClassID {
				m[i.Year]++
			}
		}
	}
	return m
}

func lessonsPerYear(db tables, year int) []string {
	var s []string
	var id, lessonid int
	var ContainsThisLesson = false
	for _, i := range db.classes.Data {
		if i.Year == year {
			id = i.ID
		}
		for _, e := range db.timetable.Data {
			if e.ClassID == id {
				lessonid = e.LessonID
			}

			for _, u := range db.lessons.Data {
				if u.ID == lessonid {
					for _, str := range s {
						if strings.Contains(str, u.Name) {
							ContainsThisLesson = true
						}
					}
					if !ContainsThisLesson {
						s = append(s, u.Name)
					}
					ContainsThisLesson = false
				}
			}
		}
	}
	return s
}

func FindStudentAndLessonID(db tables, studentname string, lessonname string) (int, int) { // to make it easier to enter data
	var studentid, lessonid int
	for _, i := range db.students.Data {
		if i.Name == studentname {
			studentid = i.ID
		}
	}
	if studentid == 0 {
		panic("StudentID is not found in the database")
	}
	for _, e := range db.lessons.Data {
		if e.Name == lessonname {
			lessonid = e.ID
		}
	}
	if lessonid == 0 {
		panic("LessonID is not found in the database")
	}
	return studentid, lessonid
}

func examsPerClass(db tables, year int, mod string) []string {
	var s []string
	var studentid, lessonid int
	var studentsid []int
	var ContainsThisExam = false
	for _, i := range db.classes.Data {
		if i.Year == year && i.Mod == mod {
			studentid = i.ID
		}
	}
	for _, e := range db.groups.Data {
		if e.ClassID == studentid {
			studentsid = append(studentsid, e.StudentID)
		}
	}
	for _, studentid := range studentsid {
		for _, u := range db.exams.Data {
			if u.StudentID == studentid {
				lessonid = u.LessonID
				for _, y := range db.lessons.Data {
					if y.ID == lessonid {
						for _, str := range s {
							if strings.Contains(str, y.Name) {
								ContainsThisExam = true
							}
						}
						if !ContainsThisExam {
							s = append(s, y.Name)
						}
						ContainsThisExam = false
					}
				}
			}
		}
	}
	return s
}

func averageGradeForStudents(db tables, name string) float64 {
	var result, lessoncount float64
	var id int
	for _, i := range db.students.Data {
		if i.Name == name {
			id = i.ID
		}
	}
	for _, e := range db.exams.Data {
		if e.StudentID == id {
			result = result + float64(e.Grade)
			lessoncount++
		}
	}
	result = result / lessoncount
	return result
}
