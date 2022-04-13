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
}

func studentCountPerClass(db tables) map[string]int {
	//it counts classes too . . .
	m := make(map[string]int)
	for _, i := range db.groups.Data {
		for _, j := range db.classes.Data {
			class := db.classes.MustFind(j.Year, j.Mod)
			if i.ClassID == j.ID {
				m[fmt.Sprint(class)]++
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

func main() {
	var db tables

	db.groups.AddStudentsToClass(
		db.classes.Insert(10, "B"),
		db.students.Insert("Jaroslavs"),
		db.students.Insert("Pavels"),
	)

	db.groups.AddStudentsToClass(
		db.classes.Insert(11, "a"),
		db.students.Insert("Abv"),
		db.students.Insert("Gd"),
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

	fmt.Println("Student count per class", studentCountPerClass(db))
	fmt.Println("Student count per year", studentCountPerYear(db))

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
