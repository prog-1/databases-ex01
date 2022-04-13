# Relational Databases

## Definition

A [relational database](https://en.wikipedia.org/wiki/Relational_database) is a digital database, where data is organizes into one or more [tables](https://en.wikipedia.org/wiki/Table_(database)) (or "relations") of [columns](https://en.wikipedia.org/wiki/Column_(database)) and [rows](https://en.wikipedia.org/wiki/Row_(database)), with a unique [key](https://en.wikipedia.org/wiki/Relational_database#Keys) identifying each row.

## Database Example

1. Table `Students`. Columns: `ID`, `Name`.

   Row examples:

   ```
   (1, "Jaroslavs")
   (2, "Pavels")
   ```

2. Table `Class`. Columns: `ID`, `Year`, `Modifier`.

   Row examples:

   ```
   (2, 10, "a")
   (3, 10, "b")
   ```

3. Table: `Groups`. Columns: `Class ID`, `Student ID`.

   Row examples:

   ```
   (3, 1)
   (3, 2)
   ```

4. Table: `Lessons`. Columns: `ID`, `Name`.

   Row examples:

   ```
   (5, "Math")
   (10, "Sport")
   (7, "Programming")
   ```

5. Table `Timetable`. Columns: `ID`, `Class ID`, `Day`, `Lesson ID`.

   Row example:

   ```
   (1, 3, "Monday", 5)
   (2, 3, "Monday", 7)
   ```
   
## Exercises

1. Implement a function `func studentCountPerClass(db tables) map[string]int` that returns the number of students in every class. The returned value is a map "class -> count", e.g. `map[string]int{"10b": 2}`.

2. Implement a function func `studentCountPerYear(db tables) map[int]int` that returns the number of students for every year. The returned value is a map "year -> count", e.g. `map[int]int{10: 2}`.

3. Implement a function `func lessonsPerYear(db tables, year int) []string` that returns a slice of unique subjects that students learn for a given year, e.g. `[]string{"Math", "Sport", "Programming"}`.

4. Create a new table `Exams` in the database with the following columns: `ID` of type `int`, `StudentID` of type `int`, `LessonID` of type `int`, `Grade` of type `int`.

   1. Write a function `func examsPerClass(db tables, year int, mod string) []string` that returns a list of exams for a given class and modifier.

   2. Write a function `func averageGradeForStudents(db tables, name string) float64` that returns an average grade for a given student.
