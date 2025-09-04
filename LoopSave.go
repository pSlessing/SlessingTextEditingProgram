package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/nsf/termbox-go"
)

func SaveAsLoop() string {
	var saveBuffer []rune

	for {
		TERMINAL.Clear()
		DisplayBuffer()
		DisplayStatus()
		PrintMessageStyle((COLS/2)-LINECOUNTWIDTH, (ROWS / 2), STYLES.MSGSTYLE, "Save As:")
		PrintMessageStyle((COLS/2)-LINECOUNTWIDTH, (ROWS/2)+1, STYLES.MSGSTYLE, string(saveBuffer))
		TERMINAL.Show()

		event := termbox.PollEvent()

		if event.Key == termbox.KeyEnter {
			filename := string(saveBuffer)
			if filename != "" {
				err := WriteBufferToFile(TEXTBUFFER, filename)
				if err != nil {
					PrintMessage(0, ROWS, tcell.ColorRed, tcell.ColorDefault,
						fmt.Sprintf("Error saving file: %s", err.Error()))
					TERMINAL.Show()
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
