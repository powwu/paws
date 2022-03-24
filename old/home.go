package main

import (
	"database/sql"
	"github.com/gdamore/tcell/v2"
	"os"
        "google.golang.org/api/classroom/v1"
	// "log"
)

func initHome(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, db *sql.DB, srv *classroom.Service, assignments [][]string) {

	start, end, current := 0, x2-x1, 0 // todo: will iterate when user presses down/up key
	// addFromClassroom(db, srv)
	drawAssignmentsList(s, x1, y1, x2-x2/5, y2, style, assignments, start, current)

out:
	for {
		s.Show()

		ev := s.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyCtrlC {
				break out // delete me later, this code only exists to prevent "out" from being unused and therefore throwing errors
				s.Fini()
				os.Exit(0)
			} else if ev.Key() == tcell.KeyCtrlR {
				addFromClassroom(db, srv)
			} else if ev.Key() == tcell.KeyEnter {
				// open url (assignments[current][3])
			} else if ev.Key() == tcell.KeyUp {
				if current <= start && start > 0 {
					start--
				} else if current <= start && start == 0 {
					break
				}
				current--
			} else if ev.Key() == tcell.KeyDown {
				if current >= end && end < len(assignments) {
					start--
				} else if current >= end && end == len(assignments)-1 {
					break
				}
				current++
			} else {
				s.Clear()
			// 	break out // "out" is the name of this cool for loop
			}
		}
	}
}

func drawAssignmentsList(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, items [][]string, start int, current int) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// Fill background
	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			s.SetContent(col, row, ' ', nil, style)
		}
	}

	// Draw borders
	for col := x1; col <= x2; col++ {
		s.SetContent(col, y1, tcell.RuneHLine, nil, style)
		s.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}

	// Only draw corners if necessary
	if y1 != y2 && x1 != x2 {
		s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
		s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
		s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
		s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
	}

	for i := y1; i < y2; i++ {
		drawText(s, x1, y1, x2/6, y2, style, items[start+y1][0])
	}


}
