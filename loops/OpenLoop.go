package loops

import "github.com/nsf/termbox-go"

func openLoop() {
	var openBuffer []rune

	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		displayBuffer()
		displayStatus()
		printMessage((COLS/2)-lineCountWidth, (ROWS / 2), termbox.ColorBlack, termbox.ColorWhite, "Open File:")
		printMessage((COLS/2)-lineCountWidth, (ROWS/2)+1, termbox.ColorBlack, termbox.ColorWhite, string(openBuffer))
		termbox.Flush()

		event := termbox.PollEvent()

		if event.Key == termbox.KeyEnter {
			openFile(string(openBuffer))
			break
		} else if event.Key == termbox.KeyBackspace || event.Key == termbox.KeyBackspace2 {
			if len(openBuffer) > 0 {
				openBuffer = openBuffer[:len(openBuffer)-1]
			}
		} else {
			openBuffer = append(openBuffer, event.Ch)
		}
	}
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	displayBuffer()
	displayStatus()
	termbox.Flush()
}
