package view

import (
	"log"

	"github.com/openmind13/gochat/message"
)

// terminal sequences
const (
	TerminalReset = "\033[0m"

	TerminalColorRed    = "\033[1;31m"
	TerminalColorYellow = "\033[1;33m"
	TerminalColorBlue   = "\033[0;34m"
)

// PrintMessage display message with colors and time
func PrintMessage(msg message.Message) {
	log.Printf("| %s%s%s > %s%s%s", TerminalColorBlue, msg.Author, TerminalReset,
		TerminalColorYellow, msg.Text, TerminalReset)
}

func PrintInfo(text string) {
	log.Printf("%s%s%s", TerminalColorRed, text, TerminalReset)
}
