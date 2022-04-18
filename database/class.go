package database

// Class represents a row storing information about a school class.
type Class struct {
	// Class ID (primary key).
	ID int
	// Class year (e.g. 1, 10, 12).
	Year int
	// Class modifier (e.g. "a", "b", "c", "d").
	Mod string
}

// ClassTable is a class database.
type ClassTable struct {
	Data []Class
}

// Insert adds a new row to the class database.
func (db *ClassTable) Insert(year int, mod string) Class {
	c := Class{
		ID:   len(db.Data) + 1,
		Year: year,
		Mod:  mod,
	}
	db.Data = append(db.Data, c)
	return c
}

// MustFind returns a class by a given year and a modifier from the class
// database.
// The class must be present in the database. Otherwise, the program panics.
func (db *ClassTable) MustFindById(id int) Class {
	for _, c := range db.Data {
		if c.ID == id {
			return c
		}
	}
	panic("not found")
}
func (db *ClassTable) MustFind(year int, mod string) Class {
	for _, c := range db.Data {
		if c.Year == year && c.Mod == mod {
			return c
		}
	}
	panic("not found")
}
