package loops

import (
	"github.com/nsf/termbox-go"
	"ste-text-editor/systemtools"
)

func WriteLoop(textBuffer [][]rune, cursorX, cursorY, offsetX, offsetY, rows, cols, lineCountWidth int) (int, int, [][]rune) {
	cursorX = lineCountWidth
	cursorY = 0
	termbox.SetCursor(cursorX, cursorY)

	for {
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey {
			switch event.Key {
			case termbox.KeyArrowUp:
				if cursorY != 0 {
					if len(textBuffer[cursorY+offsetY]) > len(textBuffer[cursorY+offsetY-1]) && cursorX-lineCountWidth > len(textBuffer[cursorY+offsetY-1]) {
						cursorX = len(textBuffer[cursorY+offsetY-1]) + lineCountWidth
						cursorY--
					} else {
						cursorY--
					}
				}
			case termbox.KeyArrowDown:
				if len(textBuffer) > cursorY+offsetY+1 {
					if len(textBuffer[cursorY+offsetY]) > len(textBuffer[cursorY+offsetY+1]) && cursorX-lineCountWidth > len(textBuffer[cursorY+offsetY+1]) {
						cursorX = len(textBuffer[cursorY+offsetY+1]) + lineCountWidth
						cursorY++
					} else {
						cursorY++
					}
				}
			case termbox.KeyArrowLeft:
				if cursorX != lineCountWidth {
					cursorX--
				}
			case termbox.KeyArrowRight:
				if cursorX-lineCountWidth < len(textBuffer[cursorY+offsetY]) {
					cursorX++
				}
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				textBuffer, cursorX, cursorY = deleteAtCursor(textBuffer, cursorX, cursorY, offsetX, offsetY, lineCountWidth)
			case termbox.KeyEnter:
				textBuffer, cursorX, cursorY = insertEnter(textBuffer, cursorX, cursorY, offsetX, offsetY, lineCountWidth)
			case termbox.KeyEsc:
				return cursorX, cursorY, textBuffer
			default:
				textBuffer, cursorX = insertRune(textBuffer, cursorX, cursorY, offsetX, offsetY, lineCountWidth, event.Ch)
			}
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

	beforeSlice := textBuffer[CursorPosYinBuffer][0:CursorPosXinBuffer]
	postSlice := textBuffer[CursorPosYinBuffer][CursorPosXinBuffer:]

	newSlice := append(beforeSlice, insertrune)
	newSlice = append(newSlice, postSlice...)
	textBuffer[CursorPosYinBuffer] = newSlice

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
