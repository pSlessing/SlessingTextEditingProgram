package main

import (
	"bufio"
	"fmt"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"os"
	"ste-text-editor/systemtools"
)

// writeBufferToFile writes the TEXTBUFFER contents to the specified file
func writeBufferToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for i, line := range TEXTBUFFER {
		// Convert rune slice to string
		lineStr := string(line)

		// Write the line
		_, err := writer.WriteString(lineStr)
		if err != nil {
			return err
		}

		// Add newline except for the last line (optional behavior)
		if i < len(TEXTBUFFER)-1 {
			_, err := writer.WriteString("\n")
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func saveCurrentState() {
	if SOURCEFILE == "" {
		// No file name set, prompt for one
		saveAsLoop()
	} else {
		// Save to existing file
		err := writeBufferToFile(SOURCEFILE)
		if err != nil {
			// Display error message to user
			systemtools.printMessage(0, ROWS, termbox.ColorRed, termbox.ColorDefault,
				fmt.Sprintf("Error saving file: %s", err.Error()))
			termbox.Flush()
		}
	}
}

// Opens a specific file and reads it into the text buffer
func openFile(filename string) {
	TEXTBUFFER = [][]rune{}
	file, err := os.Open(filename)
	if err != nil {
		SOURCEFILE = filename
		TEXTBUFFER = append(TEXTBUFFER, []rune{})
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		TEXTBUFFER = append(TEXTBUFFER, []rune{})
		count := 0
		for _, ch := range line {
			TEXTBUFFER[lineNumber] = append(TEXTBUFFER[lineNumber], rune(ch))
			count += runewidth.RuneWidth(ch)
		}
		lineNumber++
	}
	if lineNumber == 0 {
		TEXTBUFFER = append(TEXTBUFFER, []rune{})
	}
}
