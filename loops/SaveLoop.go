package loops

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"ste-text-editor/systemtools"
)

func SaveAsLoop(textBuffer [][]rune, offsetX, offsetY, rows, cols, lineCountWidth int, sourceFile string) string {
	var saveBuffer []rune

	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		systemtools.DisplayBuffer(textBuffer, offsetX, offsetY, rows, cols, lineCountWidth)
		systemtools.PrintMessage((cols/2)-lineCountWidth, (rows / 2), termbox.ColorBlack, termbox.ColorWhite, "Save As:")
		systemtools.PrintMessage((cols/2)-lineCountWidth, (rows/2)+1, termbox.ColorBlack, termbox.ColorWhite, string(saveBuffer))
		termbox.Flush()

		event := termbox.PollEvent()

		if event.Key == termbox.KeyEnter {
			filename := string(saveBuffer)
			if filename != "" {
				err := systemtools.WriteBufferToFile(textBuffer, filename)
				if err != nil {
					systemtools.PrintMessage(0, rows, termbox.ColorRed, termbox.ColorDefault,
						fmt.Sprintf("Error saving file: %s", err.Error()))
					termbox.Flush()
					termbox.PollEvent()
				} else {
					return filename
				}
			}
			break
		} else if event.Key == termbox.KeyBackspace || event.Key == termbox.KeyBackspace2 {
			if len(saveBuffer) > 0 {
				saveBuffer = saveBuffer[:len(saveBuffer)-1]
			}
		} else if event.Key == termbox.KeyEsc {
			break
		} else {
			saveBuffer = append(saveBuffer, event.Ch)
		}
	}
	return sourceFile
}
