package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
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

		event := TERMINAL.PollEvent()

		switch ev := event.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEnter {
				filename := string(saveBuffer)
				if filename != "" {
					err := WriteBufferToFile(TEXTBUFFER, filename)
					if err != nil {
						PrintMessage(0, ROWS, tcell.ColorRed, tcell.ColorDefault,
							fmt.Sprintf("Error saving file: %s", err.Error()))
						TERMINAL.Show()
						TERMINAL.PollEvent()
					} else {
						return filename
					}
				}
				return SOURCEFILE
			} else if ev.Key() == tcell.KeyBackspace || ev.Key() == tcell.KeyBackspace2 {
				if len(saveBuffer) > 0 {
					saveBuffer = saveBuffer[:len(saveBuffer)-1]
				}
			} else if ev.Key() == tcell.KeyEscape {
				return SOURCEFILE
			} else if ev.Rune() != 0 {
				saveBuffer = append(saveBuffer, ev.Rune())
			}
		}
	}
	return SOURCEFILE
}
