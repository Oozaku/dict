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

func SearchDefinition(words []string) ([]word.Word, error) {
	url := ENDPOINT + strings.Join(words, "+")
	raw, err := requests.Get(url)
	if err != nil {
		return nil, err
	}
	var entries []Entry
  err = json.Unmarshal(raw, &entries)
	if err != nil {
		log.Fatal(err)
	}
	var result []word.Word
	for _, entry := range entries {
		word := createWord(entry.Word)
		for _, phonetic := range entry.Phonetics {
			word.Phonetics = append(
				word.Phonetics,
				createPhonetic(phonetic.Phonetics, phonetic.Url),
			)
		}
		for _, meaning := range entry.Meanings {
			partOfSpeech := meaning.PartOfSpeech
			createListOfMeanings(meaning.Definitions, &word, partOfSpeech)
		}
		result = append(result, word)
	}
	return result, nil
}

func createListOfMeanings(reference []Definition, target *word.Word, partOfSpeech string) {
	for _, definition := range reference {
    newDef := word.Meaning {Definition: definition.Definition, Example: definition.UseExample}
    if val, ok := target.Meanings[partOfSpeech]; ok {
		  target.Meanings[partOfSpeech] = append(val, newDef)
    } else {
      target.Meanings[partOfSpeech] = []word.Meaning{newDef}
    }
	}
}

func createMeaning(definition, example string) word.Meaning {
	return word.Meaning{
		Definition: definition,
		Example:    example,
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
    Meanings: make(map[string][]word.Meaning),
	}
}
