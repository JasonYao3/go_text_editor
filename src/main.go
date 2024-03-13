package main

import (
	"fmt"
	"os"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

var ROWS, COLS int
var offsetX, offsetY int

var text_buffer = [][]rune{
	{'m', 'o', 'n', 'k', 'e', 'y'},
	{'e', 'd', 'i', 't', 'o', 'r'},
}

func display_text_buffer() {
	var row, col int
	for row = 0; row < ROWS; row++ {
		text_bufferRow := row + offsetY
		for col = 0; col < COLS; col++ {
			text_bufferCol := col + offsetX
			if text_bufferRow >= 0 && text_bufferRow < len(text_buffer) &&
				text_bufferCol < len(text_buffer[text_bufferRow]) {
				if text_buffer[text_bufferRow][text_bufferCol] != '\t' {
					termbox.SetChar(col, row, text_buffer[text_bufferRow][text_bufferCol])
				} else {
					termbox.SetCell(col, row, rune(' '), termbox.ColorDefault, termbox.ColorGreen)
				}
			} else if row+offsetY > len(text_buffer) {
				termbox.SetCell(0, row, rune('*'), termbox.ColorBlue, termbox.ColorDefault)
			}
		}
		termbox.SetChar(col, row, rune('\n'))
	}
}

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
		// print_message(
		// 	25,
		// 	11,
		// 	termbox.ColorDefault,
		// 	termbox.ColorDefault,
		// 	"Monkey - A bare bones text editor",
		// )
		COLS, ROWS = termbox.Size()
		ROWS--
		if COLS < 80 {
			COLS = 80
		}
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		display_text_buffer()
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
