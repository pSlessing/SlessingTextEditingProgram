package main

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

// #TODO should this be able to use any, or standard colors every time?
func PrintMessage(col, row int, fg, bg tcell.Color, msg string) {
	for _, c := range msg {
		currStyle := tcell.StyleDefault.Foreground(fg).Background(bg)
		TERMINAL.SetContent(col, row, c, nil, currStyle)
		col += runewidth.RuneWidth(c)
	}
}

func PrintMessageStyle(col, row int, style tcell.Style, msg string) {
	for _, c := range msg {
		TERMINAL.SetContent(col, row, c, nil, style)
		col += runewidth.RuneWidth(c)
	}
}

// DisplayBuffer - Pass all needed data as parameters
func DisplayBuffer() {
	var row, col int

	for row = 0; row <= ROWS; row++ {
		textBufferRow := row + OFFSETY

		DisplayLineNumber(row, textBufferRow)

		for col = 0; col < COLS; col++ {
			textBufferCol := col + OFFSETX

			if textBufferRow >= 0 &&
				textBufferRow < len(TEXTBUFFER) &&
				textBufferCol < len(TEXTBUFFER[textBufferRow]) {
				TERMINAL.SetContent(col+LINECOUNTWIDTH, row,
					TEXTBUFFER[textBufferRow][textBufferCol],
					nil, STYLES.MAINSTYLE)
			}
		}
	}
}

func DisplayStatus() {
	var col int

	for col = 0; col < COLS+LINECOUNTWIDTH; col++ {
		TERMINAL.SetContent(col, ROWS+1, ' ', nil, STYLES.STATUSSTYLE)
		if col < len(INPUTBUFFER) {
			TERMINAL.SetContent(col, ROWS+1,
				INPUTBUFFER[col],
				nil, STYLES.STATUSSTYLE)
		}
	}

	var currentLine = CURSORY + OFFSETY
	var lineNumberStr = strconv.Itoa(currentLine + 1)
	var currentColumn = CURSORX + OFFSETX - LINECOUNTWIDTH
	var columnNumberStr = strconv.Itoa(currentColumn + 1)
	// #TODO do the offsets more neat
	PrintMessageStyle(COLS, ROWS+1, STYLES.STATUSSTYLE, columnNumberStr)
	PrintMessageStyle(COLS-4, ROWS+1, STYLES.STATUSSTYLE, "col")
	PrintMessageStyle(COLS-8, ROWS+1, STYLES.STATUSSTYLE, lineNumberStr)
	PrintMessageStyle(COLS-12, ROWS+1, STYLES.STATUSSTYLE, "row")
}

func DisplayLineNumber(row int, textBufferRow int) {
	lineNumberStr := "~"

	if textBufferRow < len(TEXTBUFFER) {
		lineNumberStr = strconv.Itoa(textBufferRow + 1)
	}

	lineNumberOffset := LINECOUNTWIDTH - len(lineNumberStr)
	if lineNumberOffset > 0 {
		for i := 0; i < lineNumberOffset; i++ {
			TERMINAL.SetContent(i, row, ' ', nil, STYLES.LINECOUNTSTYLE)
		}
	}

	PrintMessageStyle(lineNumberOffset, row, STYLES.LINECOUNTSTYLE, lineNumberStr)
}

func DisplaySettingsLoop() {

}

func DisplayColorsLoop() {

}
