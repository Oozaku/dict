package main

import (
	"log"

	"github.com/Oozaku/dict/anki"
	"github.com/Oozaku/dict/config"
	"github.com/Oozaku/dict/getdef"
	"github.com/Oozaku/dict/ui"
)

// Set log to print date and location where error occurred
func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
}

func main() {

	// Get configuration file or fail program
	config, err := config.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
	}

	// Print welcome message
	ui.PrintWelcome()

	// Retrieve map of clients
	clients := getdef.RetrieveMapOfClients(config)

	for {
		// Get input from user and get list of words
		words := ui.GetEntryFromUser()

		// Call each client until get meaning without errors
		for _, client := range config.Dictionaries {

			// Get api client
			api, ok := clients[client]
			if !ok {
        log.Printf("There is no client called '%s'\n", client)
				continue
			}

			// Search for meanings
			meanings, err := api.SearchDefinition(words)
			if err != nil {
				log.Println("Could not retrieve results from", client)
				log.Println(err)
				continue
			}

			// Join words with same name
			meanings = getdef.JoinWords(meanings)

			// Print meanings and save to anki
			ui.PrintResults(meanings)
			for _, meaning := range meanings {
				err = anki.SaveWord(
					config.AnkiCsvLocation,
					meaning,
					config.AnkiMediaFolder,
				)
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
	}
}
