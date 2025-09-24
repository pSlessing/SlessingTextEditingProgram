package main

import (
	"github.com/gdamore/tcell/v2"
)

func OpenLoop() {
	var openBuffer []rune

	for {
		TERMINAL.Clear()
		DisplayBuffer()
		DisplayStatus()
		PrintMessageStyle((COLS/2)-LINECOUNTWIDTH, (ROWS / 2), STYLES.MSGSTYLE, "Open File:")
		PrintMessageStyle((COLS/2)-LINECOUNTWIDTH, (ROWS/2)+1, STYLES.MSGSTYLE, string(openBuffer))
		TERMINAL.Show()

		event := TERMINAL.PollEvent()

		switch ev := event.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEnter {
				filename := string(openBuffer)
				if filename != "" {
					newTEXTBUFFER, err := OpenFile(filename)
					if err != nil {
						// Show error but continue with current buffer
						PrintMessage(0, ROWS, tcell.ColorRed, tcell.ColorDefault, "Error opening file")
						TERMINAL.Show()
						TERMINAL.PollEvent()
						return
					}
					TEXTBUFFER = newTEXTBUFFER
					SOURCEFILE = filename
					return
				}
				break
			} else if ev.Key() == tcell.KeyBackspace || ev.Key() == tcell.KeyBackspace2 {
				if len(openBuffer) > 0 {
					openBuffer = openBuffer[:len(openBuffer)-1]
				}
			} else if ev.Key() == tcell.KeyEscape {
				break
			} else if ev.Rune() != 0 {
				openBuffer = append(openBuffer, ev.Rune())
			}
		}
	}
}
