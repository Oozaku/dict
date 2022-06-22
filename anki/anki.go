/*
Package anki contains functions to save a Word into a csv file to allow user
to import the words search in their Anki decks.

Each field in csv is separated by commas and for now, there are two fields:
front that goes in the front of card and the back, that field is on back.
*/
package anki

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Oozaku/dict/media"
	"github.com/Oozaku/dict/word"
)

// createFront creates the front of an Anki card with the word's name and the
// phonetics in audio file if possible or else in written format. The front's
// template is like bellow:
// NAME
// Phonetic1
// Phonetic2
func createFront(title string, phonetics []word.Phonetic, folder string) string {
	header := fmt.Sprintf("%s", strings.ToUpper(title))

	// Add each new non-empty phonetic in a new line
	for _, phonetic := range phonetics {

		// Add audio if possible
		if phonetic.Url != "" {

			// Get filename and save audio
			_, filename := filepath.Split(phonetic.Url)
			extension := filepath.Ext(filename)[1:]
			path := filepath.Join(folder, filename)
			err := media.DownloadAudio(phonetic.Url, path)

			// Successfully saved audio: add audio's link and go to next phonetic
			if err == nil {
				// header += fmt.Sprintf("<br/>[sound:%s]", filename)
				header += fmt.Sprintf(
					"<br/><audio controls><source src=\"%s\" type=\"audio/%s\"></audio>",
					filename,
					extension,
				)
				continue
			}

			// There was an erro: log fail and try adding phonetic's text
			log.Println(err)
		}

		// Adding audio was not possible: try adding text
		if phonetic.Text != "" {
			header += "<br/>" + phonetic.Text
		}
	}

	return header
}

// SaveWord takes the file's path and save the word inside this file, it may
// return an error if the function fails to write the file
func SaveWord(path string, word word.Word, mediaFolder string) error {

	// Open file
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create front of card, it is composed of word's name and its phonetics then
	// make it csv compliant
	front := toCsvCompliant(createFront(word.Name, word.Phonetics, mediaFolder))

	// Create card's back, composed of list of definitions and then make it csv
	// compliant
	back := toCsvCompliant(toList(word.Meanings))

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

// toList creates a list of meanings that will appear in the back of Anki's card
func toList(meanings []word.Meaning) string {
	list := ""
	for _, meaning := range meanings {
		list += fmt.Sprintf(
			"- (%s) %s<br/>",
			meaning.PartOfSpeech,
			meaning.Definition,
		)
	}
  return list
}
