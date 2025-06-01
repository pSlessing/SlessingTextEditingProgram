package loops

import (
	"github.com/nsf/termbox-go"
	"ste-text-editor/systemtools"
)

func OpenLoop(textBuffer [][]rune, offsetX, offsetY, rows, cols, lineCountWidth int, sourceFile string) ([][]rune, string) {
	var openBuffer []rune

	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		systemtools.DisplayBuffer(textBuffer, offsetX, offsetY, rows, cols, lineCountWidth)
		systemtools.PrintMessage((cols/2)-lineCountWidth, (rows / 2), termbox.ColorBlack, termbox.ColorWhite, "Open File:")
		systemtools.PrintMessage((cols/2)-lineCountWidth, (rows/2)+1, termbox.ColorBlack, termbox.ColorWhite, string(openBuffer))
		termbox.Flush()

		event := termbox.PollEvent()

		if event.Key == termbox.KeyEnter {
			filename := string(openBuffer)
			if filename != "" {
				newTextBuffer, err := systemtools.OpenFile(filename)
				if err != nil {
					// Show error but continue with current buffer
					systemtools.PrintMessage(0, rows, termbox.ColorRed, termbox.ColorDefault, "Error opening file")
					termbox.Flush()
					termbox.PollEvent()
					return textBuffer, sourceFile
				}
				return newTextBuffer, filename
			}
			break
		} else if event.Key == termbox.KeyBackspace || event.Key == termbox.KeyBackspace2 {
			if len(openBuffer) > 0 {
				openBuffer = openBuffer[:len(openBuffer)-1]
			}
		} else if event.Key == termbox.KeyEsc {
			break
		} else {
			openBuffer = append(openBuffer, event.Ch)
		}
	}
	return textBuffer, sourceFile
}
