package dictionaryapi

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/Oozaku/dict/requests"
	"github.com/Oozaku/dict/word"
)

// CONSTANTS
const ENDPOINT string = "https://api.dictionaryapi.dev/api/v2/entries/en/"

// Struct DictionaryApi holds no value
type DictionaryApi struct {}

func (dict DictionaryApi) SearchDefinition(words []string) ([]word.Word, error) {

	// Make request
	url := ENDPOINT + strings.Join(words, "+")
	headers := make(map[string]string)
	raw, err := requests.Get(url, headers)
	if err != nil {
		return nil, err
	}

	// Unmarshal response into DictionaryApi struct
	var entries []Entry
	err = json.Unmarshal(raw, &entries)
	if err != nil {
		log.Fatal(err)
	}

	// Process result and deliver list of Words as result
	var result []word.Word
	for _, entry := range entries {

		// Create a new word
		word := createWord(entry.Word)

		// Extract all valid phonetics
		for _, phonetic := range entry.Phonetics {
			word.Phonetics = append(
				word.Phonetics,
				createPhonetic(phonetic.Phonetics, phonetic.Url),
			)
		}

		// Extract meanings
		for _, meaning := range entry.Meanings {
			partOfSpeech := meaning.PartOfSpeech
			createListOfMeanings(meaning.Definitions, &word, partOfSpeech)
		}

		// Append new word
		result = append(result, word)
	}

	// Return result and nil to indicate that there was no errors
	return result, nil
}

func createListOfMeanings(reference []Definition, target *word.Word, partOfSpeech string) {
	for _, definition := range reference {
		newDef := word.Meaning{Definition: definition.Definition, PartOfSpeech: partOfSpeech}
		target.Meanings = append(target.Meanings, newDef)
	}
}

func createPhonetic(text, url string) word.Phonetic {
	return word.Phonetic{
		Text: text,
		Url:  url,
	}
}

func createWord(name string) word.Word {
	return word.Word{
		Name: name,
	}
}
