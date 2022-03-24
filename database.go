package main

import (
	_ "github.com/genjidb/genji/driver"
	"database/sql"
	"log"
	"strconv"
        "google.golang.org/api/classroom/v1"
)



func syncFromClassroom(db *sql.DB, srv *classroom.Service) { // adds all new items from classroom
	log.Print("Adding from Classroom...")
	a, err := srv.Courses.List().Do()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	for _, b := range a.Courses {
		log.Print("COURSE: " + b.Name)
		r, err := srv.Courses.CourseWork.List(b.Id).Do()
		if err != nil {
			log.Fatalf("Unable to retrieve courses. %v", err)
		}

		for _, c := range r.CourseWork {
			present := false

			resp, err := db.Query("SELECT * FROM Assignments WHERE id = ?", c.Id)
			defer resp.Close()
			if err != nil {
				log.Fatal(err)
			}



			for resp.Next() {
				log.Print("skipping " + c.Title)
				present = true // if item exists, it is present
			}

			if present == false {
				log.Print("adding " + c.Title)
				var days int64
				var date string
				if c.DueDate != nil {
					days, date = getDateInfo(c)
				}

				_, err = db.Exec(`INSERT INTO Assignments (classid, classname, id, name, url, due, days, aspen, grade, ignored, priority) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, b.Id, b.Name, c.Id, c.Title, c.AlternateLink, date, false, days, "", 0, false, 0.0)
				if err != nil {
					log.Fatalf("%+v", err)
				}
			}
		}
	}
}

func ignoreCompleted(db *sql.DB, srv *classroom.Service) {
	log.Print("Scanning for completed...")
	var ignores []string
	resp, err := db.Query("SELECT id, classid FROM Assignments WHERE ignored = false")
	if err != nil {
		log.Fatalf("%+v", err)
	}

	for resp.Next() {
		var id, classid string

		resp.Scan(&id, &classid)

		log.Print("scanning " + id)
		d, err := srv.Courses.CourseWork.StudentSubmissions.List(classid, id).Do()
		if err != nil {
			log.Fatalf("%+v", err)
		}

		for _, e := range d.StudentSubmissions {
			if e.State == "TURNED_IN" || e.State == "RETURNED" {
				ignores = append(ignores, id)
			}
		}
	}
	submitCompletedToDatabase(db, ignores)
}

func submitCompletedToDatabase(db *sql.DB, ids []string) {
	for _, id := range ids {
		ignoreEntry(db, id)
	}
}

func ignoreEntry(db *sql.DB, id string) {
	log.Print("IGNORING: " + id)
	_, err := db.Exec(`UPDATE Assignments set ignored = true WHERE id = (?)`, id)
	if err != nil {
		log.Fatalf("%+v", err)
	}
}

// func ignoreFromWeb(db *sql.DB, ps httprouter.Params) {
// 	ignoreEntry(db, ps.ByName(id))
// }


func getDateInfo(c *classroom.CourseWork) (int64, string) { // days, formatted date
	var days int64
	var date string

	if c.DueDate.Year != 0 {
		days += (c.DueDate.Year) * 365.0
		date += strconv.FormatInt(c.DueDate.Year, 10)
	}
	if c.DueDate.Month != 0 {
		days += (c.DueDate.Month) * 30.0 // yeah yeah whatever
		date += "-" + strconv.FormatInt(c.DueDate.Month, 10)
	}
	if c.DueDate.Day != 0 {
		days += c.DueDate.Day
		date += "-" + strconv.FormatInt(c.DueDate.Day - 1, 10)
	}
	return days, date
}
