package systemtools

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"ste-text-editor/systemtools"
	"strconv"
)

// Print a string starting at a specific row and column on the screen
func printMessage(col, row int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(col, row, c, fg, bg)
		col += runewidth.RuneWidth(c)
	}
}

func displayBuffer() {
	var row, col int

	for row = 0; row < main.ROWS+1; row++ {
		textBufferRow := row + main.OFFSETY

		displayLineNumber(row, textBufferRow)

		for col = 0; col < main.COLS; col++ {
			textBufferCol := col + main.OFFSETX

			if textBufferRow >= 0 &&
				textBufferRow < len(main.TEXTBUFFER) &&
				textBufferCol < len(main.TEXTBUFFER[textBufferRow]) {

				termbox.SetCell(col+main.LINECOUNTWIDTH, row,
					main.TEXTBUFFER[textBufferRow][textBufferCol],
					termbox.ColorDefault, termbox.ColorDefault)

			}
		}
	}
}

func displayStatus() {
	var col int

	for col = 0; col < main.COLS; col++ {
		termbox.SetCell(col, main.ROWS+1, ' ', termbox.ColorBlack, termbox.ColorWhite)
		if col < len(main.INPUTBUFFER) {
			termbox.SetCell(col, main.ROWS+1,
				main.INPUTBUFFER[col],
				termbox.ColorBlack, termbox.ColorWhite)
		}
	}
}

func displayLineNumber(row int, textBufferRow int) {

	// Display line numbers for all visible rows
	lineNumberStr := "~"
	lineNumberColor := termbox.ColorCyan
	bgColor := termbox.ColorWhite

	if textBufferRow < len(main.TEXTBUFFER) {
		lineNumberStr = strconv.Itoa(textBufferRow + 1)
		lineNumberColor = termbox.ColorCyan
		bgColor = termbox.ColorWhite
	}

	lineNumberOffset := main.LINECOUNTWIDTH - len(lineNumberStr)
	if lineNumberOffset > 0 {
		for i := 0; i < lineNumberOffset; i++ {
			termbox.SetCell(i, row, ' ', lineNumberColor, bgColor)
		}
	}

	printMessage(lineNumberOffset, row, lineNumberColor, bgColor, lineNumberStr)
}
