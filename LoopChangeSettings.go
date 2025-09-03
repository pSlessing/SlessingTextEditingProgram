package main

//TODO: Redo all of this unholy bullshit, start with the config in and out, colors being a struct was a mistake
import (
	//"fmt"
	"reflect"
	"slices"

	"github.com/nsf/termbox-go"
)

var ListOfColors = []termbox.Attribute{termbox.ColorBlack, termbox.ColorBlue, termbox.ColorCyan,
	termbox.ColorDarkGray, termbox.ColorGreen, termbox.ColorLightBlue, termbox.ColorLightCyan,
	termbox.ColorLightGray, termbox.ColorLightGreen, termbox.ColorLightMagenta, termbox.ColorLightRed,
	termbox.ColorLightYellow, termbox.ColorMagenta, termbox.ColorRed, termbox.ColorWhite, termbox.ColorYellow}

func ChangeSettingsLoop() {

	termbox.HideCursor()
	// - 1 because we use indexing to apply settings
	noOfSettings := reflect.TypeOf(Settings{}).NumField() - 1
	var currentPos int = 0
	var currentTempSettings Settings = GetCurrentSettings()

	AttToName := func(color termbox.Attribute) string {
		switch color {
		case termbox.ColorBlack:
			{
				return "Black"
			}
		case termbox.ColorBlue:
			{
				return "Blue"
			}
		case termbox.ColorCyan:
			{
				return "Cyan"
			}
		case termbox.ColorDarkGray:
			{
				return "Dark Gray"
			}
		case termbox.ColorGreen:
			{
				return "Green"
			}
		case termbox.ColorLightBlue:
			{
				return "Light Blue"
			}
		case termbox.ColorLightCyan:
			{
				return "Light Cyan"
			}
		case termbox.ColorLightGray:
			{
				return "Light Gray"
			}
		case termbox.ColorLightGreen:
			{
				return "Light Green"
			}
		case termbox.ColorLightMagenta:
			{
				return "Light Magenta"
			}
		case termbox.ColorLightRed:
			{
				return "Light Red"
			}
		case termbox.ColorLightYellow:
			{
				return "Light Yellow"
			}
		case termbox.ColorMagenta:
			{
				return "Magenta"
			}
		case termbox.ColorRed:
			{
				return "Red"
			}
		case termbox.ColorWhite:
			{
				return "White"
			}
		case termbox.ColorYellow:
			{
				return "Yellow"
			}
		default:
			return "Color not found in strings"
		}
	}

	ChangeMarkedSetting := func(mod int) {
		var currName reflect.StructField = reflect.TypeOf(Settings{}).FieldByIndex([]int{currentPos})
		var newColor termbox.Attribute
		var newIndex int
		switch currName.Name {
		case "BGColor":
			newIndex = (slices.Index(ListOfColors, currentTempSettings.BGColor) + mod) % len(ListOfColors)
			if newIndex < 0 {
				newIndex = len(ListOfColors) - 1
			}
			newColor = ListOfColors[newIndex]
			currentTempSettings.BGColor = newColor
		case "FGColor":
			newIndex = (slices.Index(ListOfColors, currentTempSettings.FGColor) + mod) % len(ListOfColors)
			if newIndex < 0 {
				newIndex = len(ListOfColors) - 1
			}
			newColor = ListOfColors[newIndex]
			currentTempSettings.FGColor = newColor
		case "StatusBGColor":
			newIndex = (slices.Index(ListOfColors, currentTempSettings.StatusBGColor) + mod) % len(ListOfColors)
			if newIndex < 0 {
				newIndex = len(ListOfColors) - 1
			}
			newColor = ListOfColors[newIndex]
			currentTempSettings.StatusBGColor = newColor
		case "StatusFGColor":
			newIndex = (slices.Index(ListOfColors, currentTempSettings.StatusFGColor) + mod) % len(ListOfColors)
			if newIndex < 0 {
				newIndex = len(ListOfColors) - 1
			}
			newColor = ListOfColors[newIndex]
			currentTempSettings.StatusFGColor = newColor
		case "MsgBGColor":
			newIndex = (slices.Index(ListOfColors, currentTempSettings.MsgBGColor) + mod) % len(ListOfColors)
			if newIndex < 0 {
				newIndex = len(ListOfColors) - 1
			}
			newColor = ListOfColors[newIndex]
			currentTempSettings.MsgBGColor = newColor
		case "MsgFGColor":
			newIndex = (slices.Index(ListOfColors, currentTempSettings.MsgFGColor) + mod) % len(ListOfColors)
			if newIndex < 0 {
				newIndex = len(ListOfColors) - 1
			}
			newColor = ListOfColors[newIndex]
			currentTempSettings.MsgFGColor = newColor
		case "LineCountBGColor":
			newIndex = (slices.Index(ListOfColors, currentTempSettings.LineCountBGColor) + mod) % len(ListOfColors)
			if newIndex < 0 {
				newIndex = len(ListOfColors) - 1
			}
			newColor = ListOfColors[newIndex]
			currentTempSettings.LineCountBGColor = newColor
		case "LineCountFGColor":
			newIndex = (slices.Index(ListOfColors, currentTempSettings.LineCountFGColor) + mod) % len(ListOfColors)
			if newIndex < 0 {
				newIndex = len(ListOfColors) - 1
			}
			newColor = ListOfColors[newIndex]
			currentTempSettings.LineCountFGColor = newColor

		}
	}

	RenderSettings := func() {
		termbox.SetCell(0, currentPos, '■', FGCOLOR, BGCOLOR)
		PrintMessage(1, 0, FGCOLOR, BGCOLOR, "Background")
		PrintMessage(1, 1, FGCOLOR, BGCOLOR, "Foreground")
		PrintMessage(1, 2, FGCOLOR, BGCOLOR, "Status BG")
		PrintMessage(1, 3, FGCOLOR, BGCOLOR, "Status FG")
		PrintMessage(1, 4, FGCOLOR, BGCOLOR, "Msg BG")
		PrintMessage(1, 5, FGCOLOR, BGCOLOR, "Msg FG")
		PrintMessage(1, 6, FGCOLOR, BGCOLOR, "LineCount BG")
		PrintMessage(1, 7, FGCOLOR, BGCOLOR, "LineCount FG")

		PrintMessage(len("Background")+2, 0, FGCOLOR, BGCOLOR, AttToName(currentTempSettings.BGColor))
		PrintMessage(len("Foreground")+2, 1, FGCOLOR, BGCOLOR, AttToName(currentTempSettings.FGColor))
		PrintMessage(len("Status BG")+2, 2, FGCOLOR, BGCOLOR, AttToName(currentTempSettings.StatusBGColor))
		PrintMessage(len("Status FG")+2, 3, FGCOLOR, BGCOLOR, AttToName(currentTempSettings.StatusFGColor))
		PrintMessage(len("Msg BG")+2, 4, FGCOLOR, BGCOLOR, AttToName(currentTempSettings.MsgBGColor))
		PrintMessage(len("Msg FG")+2, 5, FGCOLOR, BGCOLOR, AttToName(currentTempSettings.MsgFGColor))
		PrintMessage(len("LineCount BG")+2, 6, FGCOLOR, BGCOLOR, AttToName(currentTempSettings.LineCountBGColor))
		PrintMessage(len("LineCount FG")+2, 7, FGCOLOR, BGCOLOR, AttToName(currentTempSettings.LineCountFGColor))
	}
	//RenderExample := func() {}

	termbox.Clear(FGCOLOR, BGCOLOR)
	RenderSettings()
	termbox.Flush()

	for {
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey {
			switch event.Key {
			case termbox.KeyArrowDown:
				if currentPos != noOfSettings {
					currentPos++
				}
			case termbox.KeyArrowUp:
				if currentPos != 0 {
					currentPos--
				}
			case termbox.KeyArrowRight:
				ChangeMarkedSetting(1)
			case termbox.KeyArrowLeft:
				ChangeMarkedSetting(-1)
			case termbox.KeyEnter:
				SaveSettings(currentTempSettings)
			case termbox.KeyEsc:
				return
			}
		}
		ApplySettings(currentTempSettings)
		termbox.Clear(FGCOLOR, BGCOLOR)
		RenderSettings()
		termbox.Flush()
		//RenderExample()
	}

}
