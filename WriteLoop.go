package main

import (
	"github.com/nsf/termbox-go"
)

func WriteLoop() {
	CURSORX = LINECOUNTWIDTH
	CURSORY = 0
	termbox.SetCursor(CURSORX, CURSORY)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	DisplayBuffer()
	termbox.Flush()

	for {
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey {
			switch event.Key {
			case termbox.KeyArrowUp:
				if CURSORY > 0 {
					// Move cursor up within visible area
					CURSORY--
				} else if OFFSETY > 0 {
					// Scroll up when cursor is at top
					OFFSETY--
				}
				// Adjust cursor X if moving to a shorter line
				if CURSORY+OFFSETY < len(TEXTBUFFER) && CURSORX-LINECOUNTWIDTH > len(TEXTBUFFER[CURSORY+OFFSETY]) {
					CURSORX = len(TEXTBUFFER[CURSORY+OFFSETY]) + LINECOUNTWIDTH
				}
			case termbox.KeyArrowDown:
				if CURSORY < ROWS-1 && CURSORY+OFFSETY+1 < len(TEXTBUFFER) {
					// Move cursor down within visible area
					CURSORY++
				} else if OFFSETY+ROWS < len(TEXTBUFFER) {
					// Scroll down when cursor is at bottom
					OFFSETY++
				}
				// Adjust cursor X if moving to a shorter line
				if CURSORY+OFFSETY < len(TEXTBUFFER) && CURSORX-LINECOUNTWIDTH > len(TEXTBUFFER[CURSORY+OFFSETY]) {
					CURSORX = len(TEXTBUFFER[CURSORY+OFFSETY]) + LINECOUNTWIDTH
				}
			case termbox.KeyArrowLeft:
				if CURSORX > LINECOUNTWIDTH {
					CURSORX--
					// Horizontal scroll left if needed
					if CURSORX < LINECOUNTWIDTH {
						CURSORX = LINECOUNTWIDTH
					}
				} else if OFFSETX > 0 {
					OFFSETX--
				}
			case termbox.KeyArrowRight:
				if CURSORY+OFFSETY < len(TEXTBUFFER) {
					// Only allow moving right if not past end of line
					lineLen := len(TEXTBUFFER[CURSORY+OFFSETY])
					if CURSORX-LINECOUNTWIDTH+OFFSETX < lineLen {
						CURSORX++
						// Horizontal scroll right if needed
						if CURSORX >= COLS+LINECOUNTWIDTH {
							OFFSETX++
							CURSORX = COLS + LINECOUNTWIDTH - 1
						}
					}
				}
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				// If at the left edge and more to the left, scroll left before deleting
				if CURSORX == LINECOUNTWIDTH && OFFSETX > 0 {
					OFFSETX--
				}
				deleteAtCursor()
				// Auto-scroll if cursor goes above visible area
				if CURSORY < 0 {
					OFFSETY += CURSORY
					CURSORY = 0
				}
				// After deletion, if no characters are visible in the current row, scroll left
				visibleRow := CURSORY + OFFSETY
				if visibleRow >= 0 && visibleRow < len(TEXTBUFFER) {
					line := TEXTBUFFER[visibleRow]
					if OFFSETX >= len(line) && OFFSETX > 0 {
						OFFSETX--
						CURSORX++
					}
				}
				// Horizontal scroll left if needed after delete
				if CURSORX < LINECOUNTWIDTH && OFFSETX > 0 {
					OFFSETX--
					CURSORX = LINECOUNTWIDTH
				}
				// If at left edge and more to the left, scroll to show next char to be deleted
				if CURSORX == LINECOUNTWIDTH && OFFSETX > 0 {
					OFFSETX--
				}
			case termbox.KeyEnter:
				insertEnter()
				// Auto-scroll if cursor goes below visible area
				if CURSORY >= ROWS {
					OFFSETY += CURSORY - ROWS + 1
					CURSORY = ROWS - 1
				}
			case termbox.KeyEsc:
				return
			default:
				insertRune(event.Ch)
				// Ensure cursor is visible after insertion (horizontal scroll)
				if CURSORX >= COLS+LINECOUNTWIDTH {
					OFFSETX++
					CURSORX = COLS + LINECOUNTWIDTH - 1
				}
				if CURSORX < LINECOUNTWIDTH {
					if OFFSETX > 0 {
						OFFSETX--
						CURSORX = LINECOUNTWIDTH
					}
				}
				// Clamp cursor to end of line after insert
				lineLen := len(TEXTBUFFER[CURSORY+OFFSETY])
				if CURSORX-LINECOUNTWIDTH+OFFSETX > lineLen {
					CURSORX = lineLen - OFFSETX + LINECOUNTWIDTH
					if CURSORX < LINECOUNTWIDTH {
						CURSORX = LINECOUNTWIDTH
					}
				}
			}
		}

		// Ensure cursor stays within bounds
		if CURSORY < 0 {
			CURSORY = 0
		}
		if CURSORY >= ROWS {
			CURSORY = ROWS - 1
		}
		if CURSORX < LINECOUNTWIDTH {
			CURSORX = LINECOUNTWIDTH
		}
		if CURSORX >= COLS+LINECOUNTWIDTH {
			CURSORX = COLS + LINECOUNTWIDTH - 1
		}

		termbox.SetCursor(CURSORX, CURSORY)
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		DisplayBuffer()
		termbox.Flush()
	}
}

func insertEnter() {
	CursorPosXinBuffer := CURSORX - LINECOUNTWIDTH + OFFSETX
	CursorPosYinBuffer := CURSORY + OFFSETY

	if CursorPosYinBuffer < 0 || CursorPosYinBuffer >= len(TEXTBUFFER) {
		return
	}

	if CursorPosXinBuffer < 0 {
		CursorPosXinBuffer = 0
	}
	if CursorPosXinBuffer > len(TEXTBUFFER[CursorPosYinBuffer]) {
		CursorPosXinBuffer = len(TEXTBUFFER[CursorPosYinBuffer])
	}

	currentLine := TEXTBUFFER[CursorPosYinBuffer]
	beforeCursor := make([]rune, CursorPosXinBuffer)
	copy(beforeCursor, currentLine[:CursorPosXinBuffer])

	afterCursor := make([]rune, len(currentLine)-CursorPosXinBuffer)
	copy(afterCursor, currentLine[CursorPosXinBuffer:])

	newTEXTBUFFER := make([][]rune, len(TEXTBUFFER)+1)

	copy(newTEXTBUFFER[:CursorPosYinBuffer], TEXTBUFFER[:CursorPosYinBuffer])

	newTEXTBUFFER[CursorPosYinBuffer] = beforeCursor
	newTEXTBUFFER[CursorPosYinBuffer+1] = afterCursor

	copy(newTEXTBUFFER[CursorPosYinBuffer+2:], TEXTBUFFER[CursorPosYinBuffer+1:])
	TEXTBUFFER = newTEXTBUFFER
	CURSORX = LINECOUNTWIDTH
	CURSORY++
	return
}

func insertRune(insertrune rune) {
	CursorPosXinBuffer := CURSORX - LINECOUNTWIDTH + OFFSETX
	CursorPosYinBuffer := CURSORY + OFFSETY

	if CursorPosYinBuffer < 0 ||
		CursorPosYinBuffer >= len(TEXTBUFFER) ||
		CursorPosXinBuffer < 0 ||
		CursorPosXinBuffer > len(TEXTBUFFER[CursorPosYinBuffer]) {
		PrintMessage(0, 0, termbox.ColorDefault, termbox.ColorRed, "INSERT WAS NOT INBOUND")
		termbox.PollEvent()
		return
	}

	line := TEXTBUFFER[CursorPosYinBuffer]
	newLine := make([]rune, len(line)+1)
	copy(newLine, line[:CursorPosXinBuffer])
	newLine[CursorPosXinBuffer] = insertrune
	copy(newLine[CursorPosXinBuffer+1:], line[CursorPosXinBuffer:])
	TEXTBUFFER[CursorPosYinBuffer] = newLine
	CURSORX++
	return
}

func deleteAtCursor() {
	CursorPosXinBuffer := CURSORX - LINECOUNTWIDTH + OFFSETX
	CursorPosYinBuffer := CURSORY + OFFSETY

	if CursorPosYinBuffer < 0 || CursorPosYinBuffer >= len(TEXTBUFFER) {
		return
	}

	if CursorPosXinBuffer <= 0 {
		if CursorPosYinBuffer > 0 {
			prevLineLength := len(TEXTBUFFER[CursorPosYinBuffer-1])

			TEXTBUFFER[CursorPosYinBuffer-1] = append(TEXTBUFFER[CursorPosYinBuffer-1], TEXTBUFFER[CursorPosYinBuffer]...)

			newTEXTBUFFER := make([][]rune, len(TEXTBUFFER)-1)
			copy(newTEXTBUFFER[:CursorPosYinBuffer], TEXTBUFFER[:CursorPosYinBuffer])
			copy(newTEXTBUFFER[CursorPosYinBuffer:], TEXTBUFFER[CursorPosYinBuffer+1:])
			TEXTBUFFER = newTEXTBUFFER
			CURSORX = prevLineLength + LINECOUNTWIDTH
			CURSORY--
			return
		}
	} else {
		if CursorPosXinBuffer <= len(TEXTBUFFER[CursorPosYinBuffer]) {
			beforeSlice := TEXTBUFFER[CursorPosYinBuffer][:CursorPosXinBuffer-1]
			afterSlice := TEXTBUFFER[CursorPosYinBuffer][CursorPosXinBuffer:]
			TEXTBUFFER[CursorPosYinBuffer] = append(beforeSlice, afterSlice...)
			CURSORX--
			return
		}
	}
	return
}
