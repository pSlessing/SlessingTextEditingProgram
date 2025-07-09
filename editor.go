package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"os"
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
	{},
}

var INPUTBUFFER []rune
var LINECOUNTWIDTH = 3

var BGCOLOR = termbox.ColorBlack
var FGCOLOR = termbox.ColorWhite
var STATUSBGCOLOR = termbox.ColorWhite
var STATUSFGCOLOR = termbox.ColorBlack
var MSGBGCOLOR = termbox.ColorWhite
var MSGFGCOLOR = termbox.ColorBlack
var LINECOUNTBGCOLOR = termbox.ColorWhite
var LINECOUNTFGCOLOR = termbox.ColorCyan

func runEditor() {
	bootErr := termbox.Init()
	if bootErr != nil {
		fmt.Println(bootErr)
		fmt.Println("Error initializing termbox. STE could not launch. Error message seen above, gl troubleshooting!")
		termbox.PollEvent()
		os.Exit(1)
	}

	loadSettings()

	titleLoop()
	mainEditorLoop()
	termbox.Close()
}

func titleLoop() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	PrintMessage(25, 11, termbox.ColorDefault, termbox.ColorDefault, "STE - Slessing Text Editor")
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

		termbox.Flush()
		DisplayBuffer()
		DisplayStatus()
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
		WriteLoop()
	case "open":
		OpenLoop()
	case "save":
		saveCurrentState()
	case "saveas":
		SOURCEFILE = SaveAsLoop()
	}
	DisplayBuffer()
	DisplayStatus()
}

// Updated saveCurrentState function using systemtools
func saveCurrentState() {
	newSourceFile, err := SaveCurrentState()
	if err != nil {
		if SOURCEFILE == "" {
			// No filename set, call save-as loop
			SOURCEFILE = SaveAsLoop()
		}
		// Error was already displayed in SaveCurrentState function
	} else {
		SOURCEFILE = newSourceFile
	}
}

// Currently only color, should be expanded
func loadSettings() {
	//Check if settings exist in current path?
	//Else, create json or ini or smth file with standard settings, prob color.json
	//This should be done in another
}
