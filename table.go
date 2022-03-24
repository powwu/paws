package main

import (
	_ "github.com/genjidb/genji/driver"
	"database/sql"
	"log"
	"os"
)
func createTable(db *sql.DB) {
	log.Print("Creating table...")
	err := os.Remove("./public/table.html")
	if err != nil {
		log.Fatalf("%+v", err)
	}

	table, err := os.OpenFile("./public/table.html", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	defer table.Close()

	table.WriteString("<table><tr><th width=\"15%\">Class</th><th width=\"100%\">Name</th><th width=\"5%\">Ignore</th><th width=\"5%\">Link</th><th width=\"15%\">Due</th></tr>")

	resp, err := db.Query("SELECT classname, name, url, due, id, priority FROM Assignments WHERE ignored = false ORDER BY priority")
	if err != nil {
		log.Fatalf("%+v", err)
	}

	for resp.Next() {
		var classname, name, url, due, id, priority string

		err = resp.Scan(&classname, &name, &url, &due, &id, &priority)

		table.WriteString("<tr><td>" + classname + "</td><td>" + name + `</td><td><form action="/ignore" method="post" enctype="multipart/form-data"><input type="hidden" id="id" name="id" value="` + id + `"><br><input type="submit" value="Ignore"></form></td>` + `<td><a target="_blank" href="` + url + `">Link</a></td><td>` + due + "</td></tr>")
	}
}
