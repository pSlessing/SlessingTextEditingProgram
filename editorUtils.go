package main

import (
	"bufio"
	"fmt"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"os"
	"strconv"
)

func handleCommand() {

}

func saveCurrentState() {

}

func writeState() {

}

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

func printMessage(col, row int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(col, row, c, fg, bg)
		col += runewidth.RuneWidth(c)
	}
}

func displayBuffer() {
	var row, col int

	for row = 0; row < ROWS; row++ {
		textBufferRow := row + offsetY

		// Display line numbers
		if textBufferRow < len(textBuffer) {
			lineNumberOffset := lineCountWidth - len(strconv.Itoa(textBufferRow+1)) - 1
			printMessage(lineNumberOffset, row,
				termbox.ColorCyan, termbox.ColorDefault, strconv.Itoa(textBufferRow+1))
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
