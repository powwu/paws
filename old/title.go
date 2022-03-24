package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/mbndr/figlet4go"
	"os"
)
func initTitle(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, title, subtitle, additionaltext string) {

	drawTitleBox(s, x1, y1, x2, y2, style, title, subtitle, additionaltext)
out:
	for {
		s.Show()

		ev := s.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyCtrlC {
				s.Fini()
				os.Exit(0)
			} else {
				s.Clear()
				break out // "out" is the name of this cool for loop
			}
		}
	}

}

func drawTitle(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, title, subtitle, additionaltext string) {
	ascii := figlet4go.NewAsciiRender()
	title_ascii, _ := ascii.Render(title)
	row := drawTextCentered(s, x1, y1, x2, y2, style, title_ascii)
	drawTextCentered(s, x1, row, x2, y2, style, subtitle)
	drawTextCentered(s, x1, y1*5/2, x2, y2, style, additionaltext)
}

func drawTitleBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, title, subtitle, additionaltext string) {

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

	drawTitle(s, x2/2, y2/3, x2-1, y2-1, style, title, subtitle, additionaltext)
}
