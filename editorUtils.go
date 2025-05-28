package main

import (
	"bufio"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"os"
	"strconv"
	"strings"
)

func inputHandling() {
	event := termbox.PollEvent()

	if event.Type == termbox.EventKey {
		if event.Key == termbox.KeyEnter {
			handleCommand()
		} else if event.Key == termbox.KeyBackspace || event.Key == termbox.KeyBackspace2 {
			if len(inputBuffer) > 0 {
				inputBuffer = inputBuffer[:len(inputBuffer)-1]
			}
		} else {
			inputBuffer = append(inputBuffer, event.Ch)

		}

	}

	if event.Type == termbox.EventKey && event.Key == termbox.KeyEsc {
		return
	}

}

func handleCommand() {
	switch strings.ToLower(string(inputBuffer)) {
	case "quit":
		termbox.Close()
		os.Exit(0)
	case "write":
		writeLoop()
	}
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

	for row = 0; row < ROWS+1; row++ {
		textBufferRow := row + offsetY

		displayLineNumber(row, textBufferRow)

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

func displayLineNumber(row int, textBufferRow int) {

	// Display line numbers for all visible rows
	lineNumberStr := "~"
	lineNumberColor := termbox.ColorCyan
	bgColor := termbox.ColorWhite

	if textBufferRow < len(textBuffer) {
		lineNumberStr = strconv.Itoa(textBufferRow + 1)
		lineNumberColor = termbox.ColorCyan
		bgColor = termbox.ColorWhite
	}

	lineNumberOffset := lineCountWidth - len(lineNumberStr)
	if lineNumberOffset > 0 {
		for i := 0; i < lineNumberOffset; i++ {
			termbox.SetCell(i, row, ' ', lineNumberColor, bgColor)
		}
	}

	printMessage(lineNumberOffset, row, lineNumberColor, bgColor, lineNumberStr)
}

func insertEnter() {
	CursorPosXinBuffer := CURSORX - lineCountWidth + offsetX
	CursorPosYinBuffer := CURSORY + offsetY

	if CursorPosYinBuffer < 0 || CursorPosYinBuffer >= len(textBuffer) {
		return
	}

	if CursorPosXinBuffer < 0 {
		CursorPosXinBuffer = 0
	}
	if CursorPosXinBuffer > len(textBuffer[CursorPosYinBuffer]) {
		CursorPosXinBuffer = len(textBuffer[CursorPosYinBuffer])
	}

	currentLine := textBuffer[CursorPosYinBuffer]
	beforeCursor := make([]rune, CursorPosXinBuffer)
	copy(beforeCursor, currentLine[:CursorPosXinBuffer])

	afterCursor := make([]rune, len(currentLine)-CursorPosXinBuffer)
	copy(afterCursor, currentLine[CursorPosXinBuffer:])

	newTextBuffer := make([][]rune, len(textBuffer)+1)

	copy(newTextBuffer[:CursorPosYinBuffer], textBuffer[:CursorPosYinBuffer])

	newTextBuffer[CursorPosYinBuffer] = beforeCursor
	newTextBuffer[CursorPosYinBuffer+1] = afterCursor

	copy(newTextBuffer[CursorPosYinBuffer+2:], textBuffer[CursorPosYinBuffer+1:])

	textBuffer = newTextBuffer

	CURSORX = lineCountWidth
	CURSORY++
}

func insertRune(insertrune rune) {
	CursorPosXinBuffer := CURSORX - lineCountWidth + offsetX
	CursorPosYinBuffer := CURSORY + offsetY

	if CursorPosYinBuffer < 0 ||
		CursorPosYinBuffer >= len(textBuffer) ||
		CursorPosXinBuffer < 0 ||
		CursorPosXinBuffer > len(textBuffer[CursorPosYinBuffer]) {
		printMessage(0, 0, termbox.ColorDefault, termbox.ColorRed, "INSERT WAS NOT INBOUND")
		termbox.PollEvent()
		return
	}

	beforeSlice := textBuffer[CursorPosYinBuffer][0:CursorPosXinBuffer]
	postSlice := textBuffer[CursorPosYinBuffer][CursorPosXinBuffer:]

	newSlice := append(beforeSlice, insertrune)
	newSlice = append(newSlice, postSlice...)
	textBuffer[CursorPosYinBuffer] = newSlice

	CURSORX++
}

func deleteAtCursor() {
	CursorPosXinBuffer := CURSORX - lineCountWidth + offsetX
	CursorPosYinBuffer := CURSORY + offsetY

	//Dont access memory we dont have access to
	if CursorPosYinBuffer < 0 || CursorPosYinBuffer >= len(textBuffer) {
		return
	}

	//If cursor is at the beginning of a line
	if CursorPosXinBuffer <= 0 {
		if CursorPosYinBuffer > 0 {
			prevLineLength := len(textBuffer[CursorPosYinBuffer-1])

			textBuffer[CursorPosYinBuffer-1] = append(textBuffer[CursorPosYinBuffer-1], textBuffer[CursorPosYinBuffer]...)

			newTextBuffer := make([][]rune, len(textBuffer)-1)
			copy(newTextBuffer[:CursorPosYinBuffer], textBuffer[:CursorPosYinBuffer])
			copy(newTextBuffer[CursorPosYinBuffer:], textBuffer[CursorPosYinBuffer+1:])
			textBuffer = newTextBuffer

			CURSORY--
			CURSORX = prevLineLength + lineCountWidth
		}
	} else {
		//Normal case
		if CursorPosXinBuffer <= len(textBuffer[CursorPosYinBuffer]) {
			beforeSlice := textBuffer[CursorPosYinBuffer][:CursorPosXinBuffer-1]
			afterSlice := textBuffer[CursorPosYinBuffer][CursorPosXinBuffer:]
			textBuffer[CursorPosYinBuffer] = append(beforeSlice, afterSlice...)

			CURSORX--
		}
	}
}
