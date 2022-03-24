package main

import (
	"net/http"
	"log"
	"database/sql"

	_ "github.com/genjidb/genji/driver"

	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
)

func main() {
	srv := classroomInit()

	db, err := sql.Open("genji", "./paws.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := echo.New()

	router.Debug = true

	router.Static("/", "./public")
	router.POST("/ignore", func(c echo.Context) error {
		id := c.FormValue("id")
		ignoreEntry(db, id)
		createTable(db)
		return c.Redirect(http.StatusMovedPermanently, "table.html")
	})
	router.POST("/refresh", func(c echo.Context) error {
		syncFromClassroom(db, srv)
		ignoreCompleted(db, srv)
		prioritize(db)
		createTable(db)
		return c.Redirect(http.StatusMovedPermanently, "index.html")
	})
	router.POST("/sync-classroom", func(c echo.Context) error {
		syncFromClassroom(db, srv)
		return c.Redirect(http.StatusMovedPermanently, "index.html")
	})
	router.POST("/prioritize", func(c echo.Context) error {
		prioritize(db)
		return c.Redirect(http.StatusMovedPermanently, "index.html")
	})
	router.POST("/ignore-completed", func(c echo.Context) error {
		ignoreCompleted(db, srv)
		return c.Redirect(http.StatusMovedPermanently, "index.html")
	})
	router.POST("/refresh-table", func(c echo.Context) error {
		createTable(db)
		return c.Redirect(http.StatusMovedPermanently, "index.html")
	})

	router.Logger.Fatal(router.Start(":7297"))

}
