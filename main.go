package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/Oozaku/dict/dictapi"
	errs "github.com/Oozaku/dict/errors"
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

		// Make request to API and get result in bytes
		rawResult, err := getdef.MakeGetRequest(words)
		noResultsErr := new(errs.NoResults)
		if errors.As(err, noResultsErr) {

      // TODO: Improve usability!
			fmt.Println("No results found!")
			continue
		}

		// Parse result into struct dictapi.Entry
		entries := parseResult(rawResult)

		// Print results beautifully
		ui.PrintResults(entries)
	}
}

// Parse result from bytes into a slice of dictapi.Entry
func parseResult(bytes []byte) []dictapi.Entry {
	var entries []dictapi.Entry
	err := json.Unmarshal(bytes, &entries)
	if err != nil {
		log.Fatal(err)
	}
	return entries
}
