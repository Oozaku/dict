/*
Package owlbot adds support to Owlbot API.
*/
package owlbot

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/Oozaku/dict/requests"
	"github.com/Oozaku/dict/word"
)

// Struct Owlbot holds the token used to authenticate requests
type Owlbot struct {
	Token string `json:"token"`
}

// Owlbot's endpoint
const ENDPOINT string = "https://owlbot.info/api/v4/dictionary/"

// Method to search for definitions in Owlbot API
func (owl Owlbot) SearchDefinition(words []string) ([]word.Word, error) {

	// Build query and head for request
	query := strings.Join(words, "+")
	url := ENDPOINT + query
	headers := make(map[string]string)
	headers["Authorization"] = fmt.Sprint("Token ", owl.Token)

	// Make GET request
	raw, err := requests.Get(url, headers)
	if err != nil {
		log.Println(err)
		return []word.Word{}, err
	}

	// Unmarshal result into string
	owlbot := Word{}
	err = json.Unmarshal(raw, &owlbot)
	if err != nil {
		log.Println(err)
		return []word.Word{}, err
	}

	// Converted results will be stored hear
	result := word.Word{
		Name:      owlbot.Name,
		Phonetics: []word.Phonetic{{Text: owlbot.Phonetic}},
	}

	// Convert Owlbot struct to internal struct
	for _, meaning := range owlbot.Meanings {
		result.Meanings = append(
			result.Meanings,
			word.Meaning{
				Definition:   meaning.Definition,
				PartOfSpeech: meaning.PartOfSpeech,
			},
		)
	}

	// Return result
	return []word.Word{result}, nil
}
