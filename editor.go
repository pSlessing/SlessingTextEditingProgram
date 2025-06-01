package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"os"
	"ste-text-editor/loops"
	"ste-text-editor/systemtools"
	"strings"
)

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
	systemtools.PrintMessage(25, 11, termbox.ColorDefault, termbox.ColorDefault, "STE - Slessing Text Editor")
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
		ROWS -= 2
		COLS -= 3
		if COLS < 78 {
			COLS = 78
		}

		systemtools.DisplayBuffer(TEXTBUFFER, OFFSETX, OFFSETY, ROWS, COLS, LINECOUNTWIDTH)
		systemtools.DisplayStatus(INPUTBUFFER, ROWS, COLS)
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

func handleCommand() {
	switch strings.ToLower(string(INPUTBUFFER)) {
	case "quit":
		termbox.Close()
		os.Exit(0)
	case "write":
		// Pass all needed variables to the write loop
		CURSORX, CURSORY, TEXTBUFFER = loops.WriteLoop(TEXTBUFFER, CURSORX, CURSORY, OFFSETX, OFFSETY, ROWS, COLS, LINECOUNTWIDTH)
	case "open":
		TEXTBUFFER, SOURCEFILE = loops.OpenLoop(TEXTBUFFER, OFFSETX, OFFSETY, ROWS, COLS, LINECOUNTWIDTH, SOURCEFILE)
	case "save":
		saveCurrentState()
	case "saveas":
		SOURCEFILE = loops.SaveAsLoop(TEXTBUFFER, OFFSETX, OFFSETY, ROWS, COLS, LINECOUNTWIDTH, SOURCEFILE)
	}
}

// Updated saveCurrentState function using systemtools
func saveCurrentState() {
	newSourceFile, err := systemtools.SaveCurrentState(TEXTBUFFER, SOURCEFILE, ROWS)
	if err != nil {
		if SOURCEFILE == "" {
			// No filename set, call save-as loop
			SOURCEFILE = loops.SaveAsLoop(TEXTBUFFER, OFFSETX, OFFSETY, ROWS, COLS, LINECOUNTWIDTH, SOURCEFILE)
		}
		// Error was already displayed in SaveCurrentState function
	} else {
		SOURCEFILE = newSourceFile
	}
}
