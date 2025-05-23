package main

import (
	"fmt"
	"os"
)

import "github.com/nsf/termbox-go"

var (
	COLS int
	ROWS int
)
var offsetX, offsetY int
var sourceFile string
var textBuffer = [][]rune{
	{'H', 'e', 'l', 'l', 'o'},
	{'w', 'o', 'r', 'l', 'd'},
}

var lineCountWidth = 4

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
	printMessage(25, 11, termbox.ColorDefault, termbox.ColorDefault, "STE - Slessing Text Editor")
	for {
		COLS, ROWS = termbox.Size()
		ROWS-- // Set current terminal size
		if COLS < 78 {
			COLS = 78
		}

		termbox.Flush()

		event := termbox.PollEvent()
		if event.Type == termbox.EventKey && event.Key == termbox.KeyEnter {
			break
		}
	}
	termbox.Flush()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
}

func mainEditorLoop() {
	for {
		COLS, ROWS = termbox.Size()
		ROWS-- // Set current terminal size
		if COLS < 78 {
			COLS = 78
		}

		displayBuffer()
		termbox.Flush()
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

		event := termbox.PollEvent()
		if event.Type == termbox.EventKey && event.Key == termbox.KeyEsc {
			return
		}
	}
}
