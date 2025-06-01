package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"os"
	"ste-text-editor/systemtools"
	"strings"
)

// saveAsLoop prompts the user for a filename and saves the file
func saveAsLoop() {
	var saveBuffer []rune

	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		systemtools.displayBuffer()
		systemtools.displayStatus()
		systemtools.printMessage((COLS/2)-LINECOUNTWIDTH, (ROWS / 2), termbox.ColorBlack, termbox.ColorWhite, "Save As:")
		systemtools.printMessage((COLS/2)-LINECOUNTWIDTH, (ROWS/2)+1, termbox.ColorBlack, termbox.ColorWhite, string(saveBuffer))
		termbox.Flush()

		event := termbox.PollEvent()

		if event.Key == termbox.KeyEnter {
			filename := string(saveBuffer)
			if filename != "" {
				SOURCEFILE = filename // Set the current file
				err := writeBufferToFile(filename)
				if err != nil {
					systemtools.printMessage(0, ROWS, termbox.ColorRed, termbox.ColorDefault,
						fmt.Sprintf("Error saving file: %s", err.Error()))
					termbox.Flush()
					termbox.PollEvent()
				}
			}
			break
		} else if event.Key == termbox.KeyBackspace || event.Key == termbox.KeyBackspace2 {
			if len(saveBuffer) > 0 {
				saveBuffer = saveBuffer[:len(saveBuffer)-1]
			}
		} else if event.Key == termbox.KeyEsc {
			break // Cancel save operation
		} else {
			saveBuffer = append(saveBuffer, event.Ch)
		}
	}

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	systemtools.displayBuffer()
	systemtools.displayStatus()
	termbox.Flush()
}

func handleCommand() {
	switch strings.ToLower(string(INPUTBUFFER)) {
	case "quit":
		termbox.Close()
		os.Exit(0)
	case "write":
		writeLoop()
	case "open":
		openLoop()
	}
}
