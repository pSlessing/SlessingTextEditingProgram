package main

import (
	"github.com/gdamore/tcell/v2"
)

func ChangeSettingsLoop() {
	exampleOffset := 5
	TERMINAL.Clear()
	currentPos := 0
	styleList := STYLES.AsSlice()
	settingsLen := len(styleList) * 2
	DisplaySettingsLoop(currentPos)
	DisplayColorsLoop(exampleOffset)
	TERMINAL.Show()

	for {
		event := TERMINAL.PollEvent()
		switch ev := event.(type) {

		case *tcell.EventKey:
			mod, key, ch := ev.Modifiers(), ev.Key(), ev.Rune()
			if mod == tcell.ModNone {
				switch key {
				case tcell.KeyEnter:
					{
						//Save
						return
					}
				case tcell.KeyEsc:
					{
						return
					}
				case tcell.KeyUp:
					{
						currentPos--
						if currentPos < 0 {
							currentPos = 0
						}
					}
				case tcell.KeyDown:
					{
						currentPos++
						if currentPos == settingsLen {
							currentPos = settingsLen - 1
						}
					}
				case tcell.KeyLeft:
					{
					}
				case tcell.KeyRight:
					{
					}
				default:
					INPUTBUFFER = append(INPUTBUFFER, ch)
				}
			} else if mod == tcell.ModCtrl {

			} else if mod == tcell.ModAlt {
			}

		}
		TERMINAL.Clear()
		DisplaySettingsLoop(currentPos)
		DisplayColorsLoop(exampleOffset)
		TERMINAL.Show()
	}

}
