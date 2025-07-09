package loops

import (
	"github.com/nsf/termbox-go"
	"ste-text-editor/systemtools"
)

func WriteLoop(textBuffer [][]rune, cursorX, cursorY, offsetX, offsetY, rows, cols, lineCountWidth int) (int, int, [][]rune) {
	cursorX = lineCountWidth
	cursorY = 0
	termbox.SetCursor(cursorX, cursorY)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	systemtools.DisplayBuffer(textBuffer, offsetX, offsetY, rows, cols, lineCountWidth)
	termbox.Flush()

	for {
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey {
			switch event.Key {
			case termbox.KeyArrowUp:
				if cursorY > 0 {
					// Move cursor up within visible area
					cursorY--
				} else if offsetY > 0 {
					// Scroll up when cursor is at top
					offsetY--
				}
				// Adjust cursor X if moving to a shorter line
				if cursorY+offsetY < len(textBuffer) && cursorX-lineCountWidth > len(textBuffer[cursorY+offsetY]) {
					cursorX = len(textBuffer[cursorY+offsetY]) + lineCountWidth
				}
			case termbox.KeyArrowDown:
				if cursorY < rows-1 && cursorY+offsetY+1 < len(textBuffer) {
					// Move cursor down within visible area
					cursorY++
				} else if offsetY+rows < len(textBuffer) {
					// Scroll down when cursor is at bottom
					offsetY++
				}
				// Adjust cursor X if moving to a shorter line
				if cursorY+offsetY < len(textBuffer) && cursorX-lineCountWidth > len(textBuffer[cursorY+offsetY]) {
					cursorX = len(textBuffer[cursorY+offsetY]) + lineCountWidth
				}
			case termbox.KeyArrowLeft:
				if cursorX > lineCountWidth {
					cursorX--
					// Horizontal scroll left if needed
					if cursorX < lineCountWidth {
						cursorX = lineCountWidth
					}
				} else if offsetX > 0 {
					offsetX--
				}
			case termbox.KeyArrowRight:
				if cursorY+offsetY < len(textBuffer) {
					// Only allow moving right if not past end of line
					lineLen := len(textBuffer[cursorY+offsetY])
					if cursorX-lineCountWidth+offsetX < lineLen {
						cursorX++
						// Horizontal scroll right if needed
						if cursorX >= cols+lineCountWidth {
							offsetX++
							cursorX = cols+lineCountWidth-1
						}
					}
				}
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				// If at the left edge and more to the left, scroll left before deleting
				if cursorX == lineCountWidth && offsetX > 0 {
					offsetX--
				}
				textBuffer, cursorX, cursorY = deleteAtCursor(textBuffer, cursorX, cursorY, offsetX, offsetY, lineCountWidth)
				// Auto-scroll if cursor goes above visible area
				if cursorY < 0 {
					offsetY += cursorY
					cursorY = 0
				}
				// After deletion, if no characters are visible in the current row, scroll left
				visibleRow := cursorY + offsetY
				if visibleRow >= 0 && visibleRow < len(textBuffer) {
					line := textBuffer[visibleRow]
					if offsetX >= len(line) && offsetX > 0 {
						offsetX--
						cursorX++;
					}
				}
				// Horizontal scroll left if needed after delete
				if cursorX < lineCountWidth && offsetX > 0 {
					offsetX--
					cursorX = lineCountWidth
				}
				// If at left edge and more to the left, scroll to show next char to be deleted
				if cursorX == lineCountWidth && offsetX > 0 {
					offsetX--
				}
			case termbox.KeyEnter:
				textBuffer, cursorX, cursorY = insertEnter(textBuffer, cursorX, cursorY, offsetX, offsetY, lineCountWidth)
				// Auto-scroll if cursor goes below visible area
				if cursorY >= rows {
					offsetY += cursorY - rows + 1
					cursorY = rows - 1
				}
			case termbox.KeyEsc:
				return cursorX, cursorY, textBuffer
			default:
				textBuffer, cursorX = insertRune(textBuffer, cursorX, cursorY, offsetX, offsetY, lineCountWidth, event.Ch)
				// Ensure cursor is visible after insertion (horizontal scroll)
				if cursorX >= cols+lineCountWidth {
					offsetX++
					cursorX = cols+lineCountWidth-1
				}
				if cursorX < lineCountWidth {
					if offsetX > 0 {
						offsetX--
						cursorX = lineCountWidth
					}
				}
				// Clamp cursor to end of line after insert
				lineLen := len(textBuffer[cursorY+offsetY])
				if cursorX-lineCountWidth+offsetX > lineLen {
					cursorX = lineLen - offsetX + lineCountWidth
					if cursorX < lineCountWidth {
						cursorX = lineCountWidth
					}
				}
			}
		}

		// Ensure cursor stays within bounds
		if cursorY < 0 {
			cursorY = 0
		}
		if cursorY >= rows {
			cursorY = rows - 1
		}
		if cursorX < lineCountWidth {
			cursorX = lineCountWidth
		}
		if cursorX >= cols+lineCountWidth {
			cursorX = cols + lineCountWidth - 1
		}

		termbox.SetCursor(cursorX, cursorY)
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		systemtools.DisplayBuffer(textBuffer, offsetX, offsetY, rows, cols, lineCountWidth)
		termbox.Flush()
	}
}

func insertEnter(textBuffer [][]rune, cursorX, cursorY, offsetX, offsetY, lineCountWidth int) ([][]rune, int, int) {
	CursorPosXinBuffer := cursorX - lineCountWidth + offsetX
	CursorPosYinBuffer := cursorY + offsetY

	if CursorPosYinBuffer < 0 || CursorPosYinBuffer >= len(textBuffer) {
		return textBuffer, cursorX, cursorY
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

	return newTextBuffer, lineCountWidth, cursorY + 1
}

func insertRune(textBuffer [][]rune, cursorX, cursorY, offsetX, offsetY, lineCountWidth int, insertrune rune) ([][]rune, int) {
	CursorPosXinBuffer := cursorX - lineCountWidth + offsetX
	CursorPosYinBuffer := cursorY + offsetY

	if CursorPosYinBuffer < 0 ||
		CursorPosYinBuffer >= len(textBuffer) ||
		CursorPosXinBuffer < 0 ||
		CursorPosXinBuffer > len(textBuffer[CursorPosYinBuffer]) {
		systemtools.PrintMessage(0, 0, termbox.ColorDefault, termbox.ColorRed, "INSERT WAS NOT INBOUND")
		termbox.PollEvent()
		return textBuffer, cursorX
	}

	line := textBuffer[CursorPosYinBuffer]
	newLine := make([]rune, len(line)+1)
	copy(newLine, line[:CursorPosXinBuffer])
	newLine[CursorPosXinBuffer] = insertrune
	copy(newLine[CursorPosXinBuffer+1:], line[CursorPosXinBuffer:])
	textBuffer[CursorPosYinBuffer] = newLine

	return textBuffer, cursorX + 1
}

func deleteAtCursor(textBuffer [][]rune, cursorX, cursorY, offsetX, offsetY, lineCountWidth int) ([][]rune, int, int) {
	CursorPosXinBuffer := cursorX - lineCountWidth + offsetX
	CursorPosYinBuffer := cursorY + offsetY

	if CursorPosYinBuffer < 0 || CursorPosYinBuffer >= len(textBuffer) {
		return textBuffer, cursorX, cursorY
	}

	if CursorPosXinBuffer <= 0 {
		if CursorPosYinBuffer > 0 {
			prevLineLength := len(textBuffer[CursorPosYinBuffer-1])

			textBuffer[CursorPosYinBuffer-1] = append(textBuffer[CursorPosYinBuffer-1], textBuffer[CursorPosYinBuffer]...)

			newTextBuffer := make([][]rune, len(textBuffer)-1)
			copy(newTextBuffer[:CursorPosYinBuffer], textBuffer[:CursorPosYinBuffer])
			copy(newTextBuffer[CursorPosYinBuffer:], textBuffer[CursorPosYinBuffer+1:])

			return newTextBuffer, prevLineLength + lineCountWidth, cursorY - 1
		}
	} else {
		if CursorPosXinBuffer <= len(textBuffer[CursorPosYinBuffer]) {
			beforeSlice := textBuffer[CursorPosYinBuffer][:CursorPosXinBuffer-1]
			afterSlice := textBuffer[CursorPosYinBuffer][CursorPosXinBuffer:]
			textBuffer[CursorPosYinBuffer] = append(beforeSlice, afterSlice...)

			return textBuffer, cursorX - 1, cursorY
		}
	}

	return textBuffer, cursorX, cursorY
}
