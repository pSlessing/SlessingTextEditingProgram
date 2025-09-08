package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gdamore/tcell/v2"
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

type StyleSet struct {
	MAINSTYLE      tcell.Style
	STATUSSTYLE    tcell.Style
	MSGSTYLE       tcell.Style
	LINECOUNTSTYLE tcell.Style
}

func (s *StyleSet) AsSlice() []tcell.Style {
	return []tcell.Style{s.MAINSTYLE, s.STATUSSTYLE, s.MSGSTYLE, s.LINECOUNTSTYLE}
}

var STYLES = &StyleSet{
	MAINSTYLE:      tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack),
	STATUSSTYLE:    tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite),
	MSGSTYLE:       tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite),
	LINECOUNTSTYLE: tcell.StyleDefault.Foreground(tcell.ColorLightBlue).Background(tcell.ColorWhite),
}
var TERMINAL, bootErr = tcell.NewScreen()

var MAXWIDTH = 78

func runEditor() {
	TERMINAL.Init()
	if bootErr != nil {
		fmt.Println(bootErr)
		fmt.Println("Error initializing termbox. STE could not launch. Error message seen above, gl troubleshooting!")
		TERMINAL.PollEvent()
		os.Exit(1)
	}

	settings, err := LoadSettings()
	if err != nil {
		fmt.Printf("Error loading settings: %v\n", err)
		return
	}
	ApplySettings(settings)

	titleLoop()
	mainEditorLoop()
}

// #TODO Make this prettier?
func titleLoop() {
	TERMINAL.Clear()
	//Print title here
	TERMINAL.Show()

	for {
		TERMINAL.HideCursor()
		event := TERMINAL.PollEvent()
		switch event.(type) {
		case *tcell.EventKey:
			return
		}
	}
}

func mainEditorLoop() {
	CURSORX = LINECOUNTWIDTH
	for {
		COLS, ROWS = TERMINAL.Size()
		//Ive forgotten why this is 2, one for buffer, but why another?
		//When 1, status bar is gone, so idk man
		ROWS -= 2
		COLS -= LINECOUNTWIDTH
		if COLS < MAXWIDTH {
			COLS = MAXWIDTH
		}

		DisplayBuffer()
		DisplayStatus()
		TERMINAL.Show()
		inputHandling()
		//TERMINAL.SetCursor(CURSORX, CURSORY)

	}
}

func inputHandling() {
	event := TERMINAL.PollEvent()

	switch ev := event.(type) {

	case *tcell.EventKey:
		mod, key, ch := ev.Modifiers(), ev.Key(), ev.Rune()
		if mod == tcell.ModNone {
			switch key {
			case tcell.KeyEnter:
				{
					handleCommand()
					INPUTBUFFER = []rune{}
				}
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				{
					if len(INPUTBUFFER) > 0 {
						INPUTBUFFER = INPUTBUFFER[:len(INPUTBUFFER)-1]
					}
				}
			case tcell.KeyEsc:
				{
					return
				}
			case tcell.KeyUp:
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
			case tcell.KeyDown:
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
			case tcell.KeyLeft:
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
			case tcell.KeyRight:
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
				INPUTBUFFER = append(INPUTBUFFER, ch)
			}
		} else if mod == tcell.ModCtrl {

		} else if mod == tcell.ModAlt {
		}

	}
}

func handleCommand() {
	switch strings.ToLower(string(INPUTBUFFER)) {
	case "quit", "q":
		os.Exit(0)
	case "write", "w":
		WriteLoop()
	case "open", "o":
		OpenLoop()
		CURSORX = LINECOUNTWIDTH
		CURSORY = 0
		//termbox.SetCursor(CURSORX, CURSORY)
	case "save", "s":
		saveCurrentState()
	case "saveas", "sa":
		SOURCEFILE = SaveAsLoop()
	case "settings", "se":
		ChangeSettingsLoop()
	}
	TERMINAL.Clear()
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
