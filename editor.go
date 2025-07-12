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

// #TODO Make this prettier?
func titleLoop() {
	termbox.Clear(FGCOLOR, BGCOLOR)
	PrintMessage(25, 11, termbox.ColorDefault, termbox.ColorDefault, "STE - Slessing Text Editor")
	termbox.Flush()

	for {
		termbox.HideCursor()
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey && event.Key == termbox.KeyEnter {
			termbox.Clear(FGCOLOR, BGCOLOR)
			termbox.Flush()
			break
		}
	}
}

func mainEditorLoop() {
	CURSORX = LINECOUNTWIDTH
	for {
		COLS, ROWS = termbox.Size()
		//Ive forgotten why this is 2, one for buffer, but why another? #TODO fck around n find out
		ROWS -= 2
		COLS -= LINECOUNTWIDTH
		//Max width #TODO change to var
		if COLS < 78 {
			COLS = 78
		}

		// #TODO fix this warcrime
		DisplayBuffer()
		DisplayStatus()
		termbox.Flush()
		inputHandling()
		termbox.SetCursor(CURSORX, CURSORY)
	}
}

func inputHandling() {
	event := termbox.PollEvent()

	if event.Type == termbox.EventKey {

		switch event.Key {
		case termbox.KeyEnter:
			{
				handleCommand()
				INPUTBUFFER = []rune{}
			}
		case termbox.KeyBackspace, termbox.KeyBackspace2:
			{
				if len(INPUTBUFFER) > 0 {
					INPUTBUFFER = INPUTBUFFER[:len(INPUTBUFFER)-1]
				}
			}
		case termbox.KeyEsc:
			{
				return
			}
		case termbox.KeyArrowUp:
			{
				if CURSORY > 0 {
					// Move cursor up within visible area
					CURSORY--
				} else if OFFSETY > 0 {
					// Scroll up when cursor is at top
					OFFSETY--
				}
				// Adjust cursor X if moving to a shorter line
				if CURSORY+OFFSETY < len(TEXTBUFFER) && CURSORX-LINECOUNTWIDTH > len(TEXTBUFFER[CURSORY+OFFSETY]) {
					CURSORX = len(TEXTBUFFER[CURSORY+OFFSETY]) + LINECOUNTWIDTH
				}
			}
		case termbox.KeyArrowDown:
			{
				if CURSORY < ROWS-1 && CURSORY+OFFSETY+1 < len(TEXTBUFFER) {
					// Move cursor down within visible area
					CURSORY++
				} else if OFFSETY+ROWS < len(TEXTBUFFER) {
					// Scroll down when cursor is at bottom
					OFFSETY++
				}
				// Adjust cursor X if moving to a shorter line
				if CURSORY+OFFSETY < len(TEXTBUFFER) && CURSORX-LINECOUNTWIDTH > len(TEXTBUFFER[CURSORY+OFFSETY]) {
					CURSORX = len(TEXTBUFFER[CURSORY+OFFSETY]) + LINECOUNTWIDTH
				}
			}
		case termbox.KeyArrowLeft:
			{
				if CURSORX > LINECOUNTWIDTH {
					CURSORX--
					// Horizontal scroll left if needed
					if CURSORX < LINECOUNTWIDTH {
						CURSORX = LINECOUNTWIDTH
					}
				} else if OFFSETX > 0 {
					OFFSETX--
				}
			}
		case termbox.KeyArrowRight:
			{
				if CURSORY+OFFSETY < len(TEXTBUFFER) {
					// Only allow moving right if not past end of line
					lineLen := len(TEXTBUFFER[CURSORY+OFFSETY])
					if CURSORX-LINECOUNTWIDTH+OFFSETX < lineLen {
						CURSORX++
						// Horizontal scroll right if needed
						if CURSORX >= COLS+LINECOUNTWIDTH {
							OFFSETX++
							CURSORX = COLS + LINECOUNTWIDTH - 1
						}
					}
				}
			}
		default:
			INPUTBUFFER = append(INPUTBUFFER, event.Ch)
		}
	}
}

func handleCommand() {
	switch strings.ToLower(string(INPUTBUFFER)) {
	case "quit":
		termbox.Close()
		os.Exit(0)
	case "write":
		WriteLoop()
	case "open":
		OpenLoop()
		CURSORX = LINECOUNTWIDTH
		CURSORY = 0
		termbox.SetCursor(CURSORX, CURSORY)
	case "save":
		saveCurrentState()
	case "saveas":
		SOURCEFILE = SaveAsLoop()
	}
	termbox.Clear(FGCOLOR, BGCOLOR)
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
