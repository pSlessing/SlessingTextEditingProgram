package main

import (
	"bufio"
	"fmt"
	"github.com/nsf/termbox-go"
	"os"
)

// WriteBufferToFile writes the textBuffer contents to the specified file
func WriteBufferToFile(textBuffer [][]rune, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for i, line := range textBuffer {
		lineStr := string(line)
		_, err := writer.WriteString(lineStr)
		if err != nil {
			return err
		}

		if i < len(textBuffer)-1 {
			_, err := writer.WriteString("\n")
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// SaveCurrentState saves the current textBuffer to the sourceFile
// If no sourceFile is set, it returns an empty string to indicate save-as is needed
func SaveCurrentState() (string, error) {
	if SOURCEFILE == "" {
		// No file name set, caller should handle save-as
		return "", fmt.Errorf("no filename set")
	} else {
		// Save to existing file
		err := WriteBufferToFile(TEXTBUFFER, SOURCEFILE)
		if err != nil {
			// Display error message to user
			PrintMessage(0, ROWS, termbox.ColorRed, termbox.ColorDefault,
				fmt.Sprintf("Error saving file: %s", err.Error()))
			termbox.Flush()
			termbox.PollEvent()
			return SOURCEFILE, err
		}
		return SOURCEFILE, nil
	}
}

// OpenFile opens a specific file and reads it into a text buffer
func OpenFile(filename string) ([][]rune, error) {
	textBuffer := [][]rune{}
	file, err := os.Open(filename)
	if err != nil {
		// Return empty buffer if file doesn't exist
		return append(textBuffer, []rune{}), err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		textBuffer = append(textBuffer, []rune{})
		for _, ch := range line {
			textBuffer[lineNumber] = append(textBuffer[lineNumber], rune(ch))
		}
		lineNumber++
	}
	if lineNumber == 0 {
		textBuffer = append(textBuffer, []rune{})
	}
	return textBuffer, nil
}
