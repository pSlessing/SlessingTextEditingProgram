package main

import (
	"fmt"
	"os"
	"ste-text-editor/systemtools"
)
import "./systemtools"
import "github.com/nsf/termbox-go"

var (
	COLS    int
	ROWS    int
	CURSORX int
	CURSORY int
)
var OFFSETY = 0
var OFFSETX = 0
var SOURCEFILE string
var TEXTBUFFER = [][]rune{
	{'H', 'e', 'l', 'l', 'o'},
	{'w', 'o', 'r', 'l', 'd'},
}

var INPUTBUFFER []rune

var LINECOUNTWIDTH = 3

func runEditor() {
	bootErr := termbox.Init()
	if bootErr != nil {
		fmt.Println(bootErr)
		fmt.Println("Error initializing termbox. STE could not launch. Error message seen above, gl troubleshooting!")
		termbox.PollEvent()
		os.Exit(1)
	}

	titleLoop()

	mainEditorLoop()

	termbox.Close()
}

func titleLoop() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	systemtools.printMessage(25, 11, termbox.ColorDefault, termbox.ColorDefault, "STE - Slessing Text Editor")
	termbox.Flush()

	for {
		termbox.HideCursor()
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey && event.Key == termbox.KeyEnter {
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			termbox.Flush()
			break
		}
	}
}

func mainEditorLoop() {
	for {
		COLS, ROWS = termbox.Size()
		ROWS -= 2 // Set current terminal size
		COLS -= 3
		if COLS < 78 {
			COLS = 78
		}

		systemtools.displayBuffer()
		systemtools.displayStatus()
		termbox.Flush()
		inputHandling()
		termbox.Flush()
	}
}

func inputHandling() {
	event := termbox.PollEvent()

	if event.Type == termbox.EventKey {
		if event.Key == termbox.KeyEnter {
			handleCommand()
			INPUTBUFFER = []rune{}
		} else if event.Key == termbox.KeyBackspace || event.Key == termbox.KeyBackspace2 {
			if len(INPUTBUFFER) > 0 {
				INPUTBUFFER = INPUTBUFFER[:len(INPUTBUFFER)-1]
			}
		} else {
			INPUTBUFFER = append(INPUTBUFFER, event.Ch)

		}

	}

	if event.Type == termbox.EventKey && event.Key == termbox.KeyEsc {
		return
	}

}
