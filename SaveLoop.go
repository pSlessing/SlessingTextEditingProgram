package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

func SaveAsLoop() string {
	var saveBuffer []rune

	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		DisplayBuffer()
		DisplayStatus()
		PrintMessage((COLS/2)-LINECOUNTWIDTH, (ROWS / 2), MSGFGCOLOR, MSGBGCOLOR, "Save As:")
		PrintMessage((COLS/2)-LINECOUNTWIDTH, (ROWS/2)+1, MSGFGCOLOR, MSGBGCOLOR, string(saveBuffer))
		termbox.Flush()

		event := termbox.PollEvent()

		if event.Key == termbox.KeyEnter {
			filename := string(saveBuffer)
			if filename != "" {
				err := WriteBufferToFile(TEXTBUFFER, filename)
				if err != nil {
					PrintMessage(0, ROWS, termbox.ColorRed, termbox.ColorDefault,
						fmt.Sprintf("Error saving file: %s", err.Error()))
					termbox.Flush()
					termbox.PollEvent()
				} else {
					return filename
				}
			}
			break
		} else if event.Key == termbox.KeyBackspace || event.Key == termbox.KeyBackspace2 {
			if len(saveBuffer) > 0 {
				saveBuffer = saveBuffer[:len(saveBuffer)-1]
			}
		} else if event.Key == termbox.KeyEsc {
			break
		} else {
			saveBuffer = append(saveBuffer, event.Ch)
		}
	}
	return SOURCEFILE
}
