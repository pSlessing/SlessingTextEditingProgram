package main

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"strconv"
)

// #TODO should this be able to use any, or standard colors every time?
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

		DisplayLineNumber(row, textBufferRow)

		for col = 0; col < COLS; col++ {
			textBufferCol := col + OFFSETY

			if textBufferRow >= 0 &&
				textBufferRow < len(TEXTBUFFER) &&
				textBufferCol < len(TEXTBUFFER[textBufferRow]) {

				termbox.SetCell(col+LINECOUNTWIDTH, row,
					TEXTBUFFER[textBufferRow][textBufferCol],
					FGCOLOR, BGCOLOR)
			}
		}
	}
}

func DisplayStatus() {
	var col int

	for col = 0; col < COLS+LINECOUNTWIDTH; col++ {
		termbox.SetCell(col, ROWS+1, ' ', STATUSFGCOLOR, STATUSBGCOLOR)
		if col < len(INPUTBUFFER) {
			termbox.SetCell(col, ROWS+1,
				INPUTBUFFER[col],
				STATUSFGCOLOR, STATUSBGCOLOR)
		}
	}
}

func DisplayLineNumber(row int, textBufferRow int) {
	lineNumberStr := "~"

	if textBufferRow < len(TEXTBUFFER) {
		lineNumberStr = strconv.Itoa(textBufferRow + 1)
	}

	lineNumberOffset := LINECOUNTWIDTH - len(lineNumberStr)
	if lineNumberOffset > 0 {
		for i := 0; i < lineNumberOffset; i++ {
			termbox.SetCell(i, row, ' ', LINECOUNTFGCOLOR, LINECOUNTBGCOLOR)
		}
	}

	PrintMessage(lineNumberOffset, row, LINECOUNTFGCOLOR, LINECOUNTBGCOLOR, lineNumberStr)
}
