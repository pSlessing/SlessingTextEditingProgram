package main

import (
	"fmt"
	"os"
)

import "github.com/nsf/termbox-go"

func runEditor() {
	bootErr := termbox.Init()
	if bootErr != nil {
		fmt.Println(bootErr)
		fmt.Println("Error initializing termbox. STE could not launch. Error message seen above, gl troubleshooting!")
		termbox.PollEvent()
		os.Exit(1)
	}

	printMessage(25, 11, termbox.ColorDefault, termbox.ColorDefault, "STE - Slessing Text Editor")
	termbox.Flush()
	termbox.PollEvent()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	//mainEditorLoop()

}

func mainEditorLoop() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
		default:

		}
	}
}
