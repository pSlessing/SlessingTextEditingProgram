package systemtools

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
func DisplayBuffer(textBuffer [][]rune, offsetX, offsetY, rows, cols, lineCountWidth int) {
	var row, col int

	for row = 0; row <= rows; row++ {
		textBufferRow := row + offsetY

		DisplayLineNumber(row, textBufferRow, lineCountWidth, textBuffer)

		for col = 0; col < cols; col++ {
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

// DisplayStatus - Pass needed data as parameters
func DisplayStatus(inputBuffer []rune, rows, cols int) {
	var col int

	for col = 0; col < cols; col++ {
		termbox.SetCell(col, rows+1, ' ', termbox.ColorBlack, termbox.ColorWhite)
		if col < len(inputBuffer) {
			termbox.SetCell(col, rows+1,
				inputBuffer[col],
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
