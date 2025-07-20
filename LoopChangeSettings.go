package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"reflect"
)

func ChangeSettingsLoop() {
	noOfSettings := reflect.TypeOf(Settings{}).NumField()
	var currentPos int
	var currentTempSettings Settings = GetCurrentSettings()
	ChangeMarkedSetting := func() {

	}

	for {
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey {
			switch event.Key {
			case termbox.KeyArrowDown:
			case termbox.KeyArrowUp:
			case termbox.KeyArrowRight:
			case termbox.KeyArrowLeft:
			case termbox.KeyEnter:
			}
		}
	}

}
