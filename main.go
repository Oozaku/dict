package main

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/Oozaku/dict/anki"
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

		// Get meanings from the chosen provider
		meanings, err := getdef.GetProvider["dictionaryapi"](words)
		treatErrors(err, words)

		// There was no error: print results
		if err == nil {
			ui.PrintResults(meanings)
			for _, meaning := range meanings {
				err = anki.SaveWord("/home/oozaku/Documents/anki/anki.csv", meaning)
        if err != nil {
          log.Fatalln(err)
        }
			}
		}
	}
}

// treatErrors check if there is an error and if it is not possible to recover,
// it will kill the program
func treatErrors(err error, words []string) {
	// No errors: return
	if err == nil {
		return
	}

	// Unknown error: kill program
	var requestError *errs.ReqError
	if !errors.As(err, &requestError) {
		log.Fatalln("Unexpected error:", err)
	}

	// No results found: print not found and return
	var reqError *errs.ReqError = err.(*errs.ReqError)
	if reqError.Code == 404 {
		fmt.Printf("Meaning of '%s' not found\n", strings.Join(words, " "))
		return
	}

	// Other error in request: log and kill
	fmt.Printf("Error while searching\n")
	fmt.Printf("Code: %d", reqError.Code)
	fmt.Printf("Body: %s", reqError.Message)
	log.Fatalln(err)
}
