package main

import (
	"log"

	"github.com/Oozaku/dict/getdef"
	"github.com/Oozaku/dict/ui"
)

// Set log to print date and location where error occurred
func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
}

func main() {
	// Print welcome message
	ui.PrintWelcome()

	for {
		// Get input from user and get list of words
		words := ui.GetEntryFromUser()

		// Get meanings from the chosen provider
		meanings, err := getdef.GetProvider["dictionaryapi"](words)
		if err != nil {
			log.Fatalln(err)
		}

		// Print results beautifully
		ui.PrintResults(meanings)
	}
}
