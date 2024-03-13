package main

import (
	"fmt"
	"os"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

func print_message(col, row int, fg, bg termbox.Attribute, message string) {
	for _, char := range message {
		termbox.SetCell(col, row, char, fg, bg)
		col += runewidth.RuneWidth(char)
	}
}

func run_editor() {
	err := termbox.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for {
		print_message(
			25,
			11,
			termbox.ColorDefault,
			termbox.ColorDefault,
			"Monkey - A bare bones text editor",
		)
		termbox.Flush()
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey && event.Key == termbox.KeyEsc {
			termbox.Close()
			break
		}
	}
}

func main() {
	run_editor()
}
