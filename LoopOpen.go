package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/nsf/termbox-go"
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

		event := termbox.PollEvent()

		if event.Key == termbox.KeyEnter {
			filename := string(openBuffer)
			if filename != "" {
				newTEXTBUFFER, err := OpenFile(filename)
				if err != nil {
					// Show error but continue with current buffer
					PrintMessage(0, ROWS, tcell.ColorRed, tcell.ColorDefault, "Error opening file")
					TERMINAL.Show()
					termbox.PollEvent()
					return
				}
				TEXTBUFFER = newTEXTBUFFER
				SOURCEFILE = filename
				return
			}
			break
		} else if event.Key == termbox.KeyBackspace || event.Key == termbox.KeyBackspace2 {
			if len(openBuffer) > 0 {
				openBuffer = openBuffer[:len(openBuffer)-1]
			}
		} else if event.Key == termbox.KeyEsc {
			break
		} else {
			openBuffer = append(openBuffer, event.Ch)
		}
	}
}
