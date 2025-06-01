package loops

import "github.com/nsf/termbox-go"

func insertEnter() {
	CursorPosXinBuffer := CURSORX - lineCountWidth + offsetX
	CursorPosYinBuffer := CURSORY + offsetY

	if CursorPosYinBuffer < 0 || CursorPosYinBuffer >= len(textBuffer) {
		return
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

	textBuffer = newTextBuffer

	CURSORX = lineCountWidth
	CURSORY++
}

func insertRune(insertrune rune) {
	CursorPosXinBuffer := CURSORX - lineCountWidth + offsetX
	CursorPosYinBuffer := CURSORY + offsetY

	if CursorPosYinBuffer < 0 ||
		CursorPosYinBuffer >= len(textBuffer) ||
		CursorPosXinBuffer < 0 ||
		CursorPosXinBuffer > len(textBuffer[CursorPosYinBuffer]) {
		printMessage(0, 0, termbox.ColorDefault, termbox.ColorRed, "INSERT WAS NOT INBOUND")
		termbox.PollEvent()
		return
	}

	beforeSlice := textBuffer[CursorPosYinBuffer][0:CursorPosXinBuffer]
	postSlice := textBuffer[CursorPosYinBuffer][CursorPosXinBuffer:]

	newSlice := append(beforeSlice, insertrune)
	newSlice = append(newSlice, postSlice...)
	textBuffer[CursorPosYinBuffer] = newSlice

	CURSORX++
}

func deleteAtCursor() {
	CursorPosXinBuffer := CURSORX - lineCountWidth + offsetX
	CursorPosYinBuffer := CURSORY + offsetY

	//Dont access memory we dont have access to
	if CursorPosYinBuffer < 0 || CursorPosYinBuffer >= len(textBuffer) {
		return
	}

	//If cursor is at the beginning of a line
	if CursorPosXinBuffer <= 0 {
		if CursorPosYinBuffer > 0 {
			prevLineLength := len(textBuffer[CursorPosYinBuffer-1])

			textBuffer[CursorPosYinBuffer-1] = append(textBuffer[CursorPosYinBuffer-1], textBuffer[CursorPosYinBuffer]...)

			newTextBuffer := make([][]rune, len(textBuffer)-1)
			copy(newTextBuffer[:CursorPosYinBuffer], textBuffer[:CursorPosYinBuffer])
			copy(newTextBuffer[CursorPosYinBuffer:], textBuffer[CursorPosYinBuffer+1:])
			textBuffer = newTextBuffer

			CURSORY--
			CURSORX = prevLineLength + lineCountWidth
		}
	} else {
		//Normal case
		if CursorPosXinBuffer <= len(textBuffer[CursorPosYinBuffer]) {
			beforeSlice := textBuffer[CursorPosYinBuffer][:CursorPosXinBuffer-1]
			afterSlice := textBuffer[CursorPosYinBuffer][CursorPosXinBuffer:]
			textBuffer[CursorPosYinBuffer] = append(beforeSlice, afterSlice...)

			CURSORX--
		}
	}
}

func writeLoop() {
	CURSORX = lineCountWidth
	CURSORY = 0
	termbox.SetCursor(CURSORX, CURSORY)
	for {
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey {
			switch event.Key {
			case termbox.KeyArrowUp:
				if CURSORY != 0 {
					//If the length of the current line is more than the next
					if len(textBuffer[CURSORY+offsetY]) > len(textBuffer[CURSORY+offsetY-1]) && CURSORX-lineCountWidth > len(textBuffer[CURSORY+offsetY-1]) {
						CURSORX = len(textBuffer[CURSORY+offsetY-1]) + lineCountWidth
						CURSORY--
					} else {
						CURSORY--
					}
				}
			case termbox.KeyArrowDown:
				//Is there a line below?
				if len(textBuffer) > CURSORY+offsetY+1 {
					//If the length of the current line is more than the next
					if len(textBuffer[CURSORY+offsetY]) > len(textBuffer[CURSORY+offsetY+1]) && CURSORX-lineCountWidth > len(textBuffer[CURSORY+offsetY+1]) {
						CURSORX = len(textBuffer[CURSORY+offsetY+1]) + lineCountWidth
						CURSORY++
					} else {
						CURSORY++
					}
				}
			case termbox.KeyArrowLeft:
				if CURSORX != lineCountWidth {
					CURSORX--
				}
			case termbox.KeyArrowRight:
				if CURSORX-lineCountWidth < len(textBuffer[CURSORY+offsetY]) {
					CURSORX++
				}
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				deleteAtCursor()
			case termbox.KeyEnter:
				insertEnter()
			default:
				insertRune(event.Ch)
			}
		}

		termbox.SetCursor(CURSORX, CURSORY)
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		displayBuffer()
		termbox.Flush()
	}
}
