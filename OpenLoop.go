package main

import (
	"github.com/nsf/termbox-go"
)

func OpenLoop() {
	var openBuffer []rune

	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		DisplayBuffer()
		PrintMessage((COLS/2)-LINECOUNTWIDTH, (ROWS / 2), MSGFGCOLOR, MSGBGCOLOR, "Open File:")
		PrintMessage((COLS/2)-LINECOUNTWIDTH, (ROWS/2)+1, MSGFGCOLOR, MSGBGCOLOR, string(openBuffer))
		termbox.Flush()

		event := termbox.PollEvent()

		if event.Key == termbox.KeyEnter {
			filename := string(openBuffer)
			if filename != "" {
				newTEXTBUFFER, err := OpenFile(filename)
				if err != nil {
					// Show error but continue with current buffer
					PrintMessage(0, ROWS, termbox.ColorRed, termbox.ColorDefault, "Error opening file")
					termbox.Flush()
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
