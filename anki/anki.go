/*
Package anki contains functions to save a Word into a csv file to allow user
to import the words search in their Anki decks.

Each field in csv is separated by commas and for now, there are two fields:
front that goes in the front of card and the back, that field is on back.
*/
package anki

import (
	"fmt"
	"os"
	"strings"

	"github.com/Oozaku/dict/word"
)

// createFront creates the front of an Anki card with the word's name and the
// phonetics in Html format. The format is something like:
// NAME
// Phonetic1
// Phonetic2
func createFront(title string, phonetics []word.Phonetic) string {
	header := fmt.Sprintf("%s", strings.ToUpper(title))

	// Add each new non-empty phonetic in a new line
	for _, phonetic := range phonetics {
		if phonetic.Text != "" {
			header += "<br/>" + phonetic.Text
		}
	}

	return header
}

// SaveWord takes the file's path and save the word inside this file, it may
// return an error if the function fails to write the file
func SaveWord(path string, word word.Word) error {

	// Open file
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create front of card, it is composed of word's name and its phonetics then
	// make it csv compliant
	front := toCsvCompliant(createFront(word.Name, word.Phonetics))

	// Create card's back, composed of list of definitions and then make it csv
	// compliant
	back := toCsvCompliant(toHtmlList(word.Meanings))

	// Add new line in csv, each field separated by comma
	_, err = fmt.Fprintf(file, "%s, %s\n", front, back)
	if err != nil {
		return err
	}

	return nil
}

// toCsvCompliant surrounds the field inside double quotes and escapes
// each double quotes inside field with another double quotes as specified in
// RFC 4180
func toCsvCompliant(field string) string {
	return fmt.Sprintf("\"%s\"", strings.ReplaceAll(field, "\"", "\"\""))
}

// toHtmlList transforms a map of meanings grouped by the part of speech into a
// Html ordered list. The format that this function follows is:
// <number>. (part of speech) meaning
func toHtmlList(meanings map[string][]word.Meaning) string {
	list := "<ol>"
	for partOfSpeech, definitions := range meanings {
		for _, definition := range definitions {
			list += fmt.Sprintf("<li>(%s) %s</li>", partOfSpeech, definition.Definition)
		}
	}
	return list + "</ol>"
}
