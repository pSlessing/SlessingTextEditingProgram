package main

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"strconv"
)

// PrintMessage - Exported function (capital P)
func PrintMessage(col, row int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(col, row, c, fg, bg)
		col += runewidth.RuneWidth(c)
	}
}

// DisplayBuffer - Pass all needed data as parameters
func DisplayBuffer() {
	var row, col int

	for row = 0; row <= ROWS; row++ {
		textBufferRow := row + OFFSETY

		DisplayLineNumber(row, textBufferRow, LINECOUNTWIDTH, TEXTBUFFER)

		for col = 0; col < COLS; col++ {
			textBufferCol := col + OFFSETY

			if textBufferRow >= 0 &&
				textBufferRow < len(TEXTBUFFER) &&
				textBufferCol < len(TEXTBUFFER[textBufferRow]) {

				termbox.SetCell(col+LINECOUNTWIDTH, row,
					TEXTBUFFER[textBufferRow][textBufferCol],
					termbox.ColorDefault, termbox.ColorDefault)
			}
		}
	}
}

// DisplayStatus - Pass needed data as parameters
func DisplayStatus() {
	var col int

	for col = 0; col < COLS+LINECOUNTWIDTH; col++ {
		termbox.SetCell(col, ROWS+1, ' ', termbox.ColorBlack, termbox.ColorWhite)
		if col < len(INPUTBUFFER) {
			termbox.SetCell(col, ROWS+1,
				INPUTBUFFER[col],
				termbox.ColorBlack, termbox.ColorWhite)
		}
	}
}

// DisplayLineNumber - Helper function
func DisplayLineNumber(row int, textBufferRow int, lineCountWidth int, textBuffer [][]rune) {
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

	PrintMessage(lineNumberOffset, row, lineNumberColor, bgColor, lineNumberStr)
}
