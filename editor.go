package main

import (
	"fmt"
	"os"
)

import "github.com/nsf/termbox-go"

var (
	COLS    int
	ROWS    int
	CURSORX int
	CURSORY int
)
var offsetY = 0
var offsetX = 0
var sourceFile string
var textBuffer = [][]rune{
	{'H', 'e', 'l', 'l', 'o'},
	{'w', 'o', 'r', 'l', 'd'},
}

var inputBuffer []rune

var lineCountWidth = 3

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
	printMessage(25, 11, termbox.ColorDefault, termbox.ColorDefault, "STE - Slessing Text Editor")
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

		displayBuffer()
		displayStatus()
		termbox.Flush()
		inputHandling()
		termbox.Flush()
	}
}

func writeLoop() {
	CURSORX = lineCountWidth
	CURSORY = 0
	termbox.SetCursor(CURSORX, CURSORY)
	for {
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey {
			switch event.Key {
			case termbox.KeyArrowUp:
				if CURSORY != 0 {
					CURSORY--
				}
			case termbox.KeyArrowDown:
				if len(textBuffer) > CURSORY+offsetY+1 {
					if len(textBuffer[CURSORY+offsetY]) > len(textBuffer[CURSORY+offsetY+1]) {
						CURSORY = len(textBuffer[CURSORY+offsetY+1])
					} else {
						CURSORY++
					}
				}
			case termbox.KeyArrowLeft:
				if CURSORX != lineCountWidth {
					CURSORX--
				}
			case termbox.KeyArrowRight:
				if CURSORX-lineCountWidth < len(textBuffer[CURSORY+offsetY]) {
					CURSORX++
				}
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				deleteAtCursor()
			case termbox.KeyEnter:
				insertEnter()
			default:
				insertRune(event.Ch)
			}
		}

		termbox.SetCursor(CURSORX, CURSORY)
		displayBuffer()
		termbox.Flush()
	}
}
