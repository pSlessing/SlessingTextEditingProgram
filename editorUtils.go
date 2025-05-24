package main

import (
	"bufio"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"os"
	"strconv"
)

func inputHandling() {
	event := termbox.PollEvent()

	if event.Type == termbox.EventKey {
		inputBuffer = append(inputBuffer, event.Ch)
	}

	if event.Type == termbox.EventKey && event.Key == termbox.KeyEsc {
		return
	}

}

func handleCommand() {

}

func saveCurrentState() {

}

func writeState() {

}

// Opens a specific file and reads it into the text buffer
func openFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		sourceFile = filename
		textBuffer = append(textBuffer, []rune{})
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		textBuffer = append(textBuffer, []rune{})
		count := 0
		for _, ch := range line {
			textBuffer[lineNumber] = append(textBuffer[lineNumber], rune(ch))
			count += runewidth.RuneWidth(ch)
		}
		lineNumber++
	}
	if lineNumber == 0 {
		textBuffer = append(textBuffer, []rune{})
	}
}

// Print a string starting at a specific row and column on the screen
func printMessage(col, row int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(col, row, c, fg, bg)
		col += runewidth.RuneWidth(c)
	}
}

// Stupid way to do this, everything in one function, no options for modularity
// One could maybe include a list of lists of functions, each index being ran at a specific point, maybe.
// Would potentially cause trouble at certain points (for example, how does one just insert a row option wihout needing new variables in the base function)
func displayBuffer() {
	var row, col int

	for row = 0; row < ROWS; row++ {
		textBufferRow := row + offsetY

		// Display line numbers
		if textBufferRow < len(textBuffer) {
			lineNumberOffset := lineCountWidth - len(strconv.Itoa(textBufferRow+1)) - 1
			printMessage(lineNumberOffset, row,
				termbox.ColorCyan, termbox.ColorWhite, strconv.Itoa(textBufferRow+1))
		}

		for col = 0; col < COLS; col++ {
			textBufferCol := col + offsetX

			if textBufferRow >= 0 &&
				textBufferRow < len(textBuffer) &&
				textBufferCol < len(textBuffer[textBufferRow]) {

				termbox.SetCell(col+lineCountWidth, row,
					textBuffer[textBufferRow][textBufferCol],
					termbox.ColorDefault, termbox.ColorDefault)

			}
		}
	}
}

func displayStatus() {
	var col int

	for col = 0; col < COLS; col++ {
		termbox.SetCell(col, ROWS+1, ' ', termbox.ColorBlack, termbox.ColorWhite)
		if col < len(inputBuffer) {
			termbox.SetCell(col, ROWS+1,
				inputBuffer[col],
				termbox.ColorBlack, termbox.ColorWhite)
		}
	}
}
