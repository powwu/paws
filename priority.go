// priority:
// - start at priority 0
// - for every day until due date, add one priority
// - for grade percentage in class, add equivalent priority (will need to tie classes to aspen, which is fine)

package main

import (
	"database/sql"
	"time"
	"log"
)

func prioritize(db *sql.DB) {
	resp, err := db.Query("SELECT priority, days, grade, id, aspen FROM Assignments")
	if err != nil {
	 	log.Fatalf("%+v", err)
	}
	defer resp.Close()

	m := make(map[string]int64)

	for resp.Next() {
		var priority, days, grade int64
		var id, aspen string

		currentday := 0
		year, month, day := time.Now().Date()

		resp.Scan(&priority, &days, &grade, &id, &aspen)

		log.Print("prioritizing: " + id)

		currentday += int(year) * 365
		currentday += int(month) * 30
		currentday += int(day)

		priority += days - int64(currentday)

		priority += grade

		m[id] = priority
	}
	submitPriorityToDatabase(db, m)
}

func submitPriorityToDatabase(db *sql.DB, m map[string]int64) {
	for id, priority := range m {
		_, err := db.Exec(`UPDATE Assignments SET priority = (?) WHERE id = (?)`, priority, id)
		if err != nil {
			log.Fatalf("%+v", err)
		}
	}
}
