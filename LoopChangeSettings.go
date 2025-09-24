package main

import (
	"github.com/gdamore/tcell/v2"
)

func ChangeSettingsLoop() {

	colors := map[string]tcell.Color{
		// Basic colors
		"Black":   tcell.ColorBlack,
		"Red":     tcell.ColorRed,
		"Green":   tcell.ColorGreen,
		"Yellow":  tcell.ColorYellow,
		"Blue":    tcell.ColorBlue,
		"Magenta": tcell.ColorDarkMagenta,
		"Cyan":    tcell.ColorDarkCyan,
		"White":   tcell.ColorWhite,

		// Extended colors
		"Gray":     tcell.ColorGray,
		"DarkGray": tcell.ColorDarkGray,
		"Silver":   tcell.ColorSilver,
		"Maroon":   tcell.ColorMaroon,
		"Olive":    tcell.ColorOlive,
		"Lime":     tcell.ColorLime,
		"Aqua":     tcell.ColorAqua,
		"Teal":     tcell.ColorTeal,
		"Navy":     tcell.ColorNavy,
		"Fuchsia":  tcell.ColorFuchsia,
		"Purple":   tcell.ColorPurple,
		"Orange":   tcell.ColorOrange,
	}

	colorNames := []string{
		"Black", "Red", "Green", "Yellow", "Blue", "Magenta", "Cyan", "White",
		"Gray", "DarkGray", "Silver", "Maroon", "Olive", "Lime", "Aqua", "Teal",
		"Navy", "Fuchsia", "Purple", "Orange",
	}

	exampleOffset := 5
	TERMINAL.Clear()
	currentPos := 0
	styleList := STYLES.AsSlice()
	settingsLen := len(styleList) * 2
	colorPos := 0

	// Initialize colorPos to match the current setting's color
	colorPos = getCurrentColorPos(currentPos, colorNames)

	DisplaySettingsLoop(currentPos)
	DisplayColorsLoop(exampleOffset)
	TERMINAL.Show()

	for {
		event := TERMINAL.PollEvent()
		switch ev := event.(type) {

		case *tcell.EventKey:
			mod, key := ev.Modifiers(), ev.Key()
			if mod == tcell.ModNone {
				switch key {
				case tcell.KeyEnter:
					{
						// Apply selected color to current style setting
						if currentPos < settingsLen {
							selectedColor := colors[colorNames[colorPos]]
							styleIndex := currentPos / 2
							isBackground := (currentPos % 2) == 0
							updateStylesHelper(styleIndex, selectedColor, isBackground)

						}
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

						// Update colorPos to match the current setting's color
						if currentPos < settingsLen {
							colorPos = getCurrentColorPos(currentPos, colorNames)
						}
					}
				case tcell.KeyDown:
					{
						currentPos++
						if currentPos == settingsLen {
							currentPos = settingsLen - 1
						}

						// Update colorPos to match the current setting's color
						if currentPos < settingsLen {
							colorPos = getCurrentColorPos(currentPos, colorNames)
						}
					}
				case tcell.KeyLeft:
					{
						// Navigate through colors
						colorPos--
						if colorPos < 0 {
							colorPos = len(colorNames) - 1
						}

						// Apply selected color immediately for preview
						if currentPos < settingsLen {
							selectedColor := colors[colorNames[colorPos]]
							styleIndex := currentPos / 2
							isBackground := (currentPos % 2) == 0
							updateStylesHelper(styleIndex, selectedColor, isBackground)
						}
					}
				case tcell.KeyRight:
					{
						// Navigate through colors
						colorPos++
						if colorPos >= len(colorNames) {
							colorPos = 0
						}

						// Apply selected color immediately for preview
						if currentPos < settingsLen {
							selectedColor := colors[colorNames[colorPos]]
							styleIndex := currentPos / 2
							isBackground := (currentPos % 2) == 0
							updateStylesHelper(styleIndex, selectedColor, isBackground)
						}
					}
				default:
				}
			} else if mod == tcell.ModCtrl {

			} else if mod == tcell.ModAlt {
			}

		}
		TERMINAL.Clear()

		// Save the updated settings
		currentSettings := GetCurrentSettings()
		err := SaveSettings(currentSettings)
		if err != nil {
			// Show error message
			PrintMessage(0, ROWS-1, tcell.ColorRed, tcell.ColorDefault, "Error saving settings")
			TERMINAL.Show()
			TERMINAL.PollEvent() // Wait for user input
		}

		// Pass current color selection to display functions for live preview
		DisplaySettingsLoop(currentPos)
		DisplayColorsLoop(exampleOffset)

		TERMINAL.Show()
	}

}

// Helper function to get the current color position based on the selected setting
func getCurrentColorPos(currentPos int, colorNames []string) int {
	if currentPos >= len(STYLES.AsSlice())*2 {
		return 0
	}

	styleIndex := currentPos / 2
	isBackground := (currentPos % 2) == 0

	var currentColor tcell.Color
	styleList := STYLES.AsSlice()

	if isBackground {
		_, currentColor, _ = styleList[styleIndex].Decompose()
	} else {
		currentColor, _, _ = styleList[styleIndex].Decompose()
	}

	// Find the color name that matches the current color
	for i, colorName := range colorNames {
		if getColorFromName(colorName) == currentColor {
			return i
		}
	}

	// If color not found in our list, return 0 (default to first color)
	return 0
}

// Helper function to convert color name to tcell.Color
func getColorFromName(colorName string) tcell.Color {
	colors := map[string]tcell.Color{
		"Black":    tcell.ColorBlack,
		"Red":      tcell.ColorRed,
		"Green":    tcell.ColorGreen,
		"Yellow":   tcell.ColorYellow,
		"Blue":     tcell.ColorBlue,
		"Magenta":  tcell.ColorDarkMagenta,
		"Cyan":     tcell.ColorDarkCyan,
		"White":    tcell.ColorWhite,
		"Gray":     tcell.ColorGray,
		"DarkGray": tcell.ColorDarkGray,
		"Silver":   tcell.ColorSilver,
		"Maroon":   tcell.ColorMaroon,
		"Olive":    tcell.ColorOlive,
		"Lime":     tcell.ColorLime,
		"Aqua":     tcell.ColorAqua,
		"Teal":     tcell.ColorTeal,
		"Navy":     tcell.ColorNavy,
		"Fuchsia":  tcell.ColorFuchsia,
		"Purple":   tcell.ColorPurple,
		"Orange":   tcell.ColorOrange,
	}

	if color, exists := colors[colorName]; exists {
		return color
	}
	return tcell.ColorDefault
}

func updateStylesHelper(switchIndex int, selectedColor tcell.Color, isBackground bool) {
	switch switchIndex {
	case 0: // Main style
		if isBackground {
			fg, _, _ := STYLES.MAINSTYLE.Decompose()
			STYLES.MAINSTYLE = tcell.StyleDefault.Background(selectedColor).Foreground(fg)
		} else {
			_, bg, _ := STYLES.MAINSTYLE.Decompose()
			STYLES.MAINSTYLE = tcell.StyleDefault.Background(bg).Foreground(selectedColor)
		}
	case 1: // Status style
		if isBackground {
			fg, _, _ := STYLES.STATUSSTYLE.Decompose()
			STYLES.STATUSSTYLE = tcell.StyleDefault.Background(selectedColor).Foreground(fg)
		} else {
			_, bg, _ := STYLES.STATUSSTYLE.Decompose()
			STYLES.STATUSSTYLE = tcell.StyleDefault.Background(bg).Foreground(selectedColor)
		}
	case 2: // Message style
		if isBackground {
			fg, _, _ := STYLES.MSGSTYLE.Decompose()
			STYLES.MSGSTYLE = tcell.StyleDefault.Background(selectedColor).Foreground(fg)
		} else {
			_, bg, _ := STYLES.MSGSTYLE.Decompose()
			STYLES.MSGSTYLE = tcell.StyleDefault.Background(bg).Foreground(selectedColor)
		}
	case 3: // Line count style
		if isBackground {
			fg, _, _ := STYLES.LINECOUNTSTYLE.Decompose()
			STYLES.LINECOUNTSTYLE = tcell.StyleDefault.Background(selectedColor).Foreground(fg)
		} else {
			_, bg, _ := STYLES.LINECOUNTSTYLE.Decompose()
			STYLES.LINECOUNTSTYLE = tcell.StyleDefault.Background(bg).Foreground(selectedColor)
		}
	}
}
