package main

import (
	"log"
	"os"
	_ "github.com/genjidb/genji/driver"
	"database/sql"

	"github.com/gdamore/tcell/v2"
)
func main() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple)
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	screen.Init()
	w, h := screen.Size()
	screen.SetStyle(defStyle)
	screen.EnableMouse()
	screen.EnablePaste()

	srv := classroomInit()

	db, err := sql.Open("genji", "./paws.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var assignments [][]string = getAllFromDatabase(db)

	initTitle(screen, 1, 1, w-1, h-1, boxStyle, "PAWS", "Powwu's Academic Work Suite", "Press any key to continue")



	initHome(screen, 1, 1, w-1, h-1, defStyle, db, srv, assignments)

	for {
		screen.Show()

		ev := screen.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			screen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyCtrlC {
				screen.Fini()
				os.Exit(0)
		}
	}

}

}
