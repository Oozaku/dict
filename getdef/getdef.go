package getdef

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/Oozaku/dict/errors"
)

// Endpoint to API
const URL string = "https://api.dictionaryapi.dev/api/v2/entries/en/"

// Make a get request to API and return the raw result
func MakeGetRequest(words []string) ([]byte, error) {

	// Join words to make query
	url := makeQueryURL(words)

	// Get data from API
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	// Close connection at the end
	defer resp.Body.Close()

	// Response returns not found: raise error
	if resp.StatusCode == 404 {
		return []byte{}, errors.NoResults{ErrorMessage: "No results found", ErrorCode: 404}
	}

	// Read all data from response and return it
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body, nil
}

// Join slice of words with +
func makeQueryURL(words []string) string {
	query := strings.Join(words, "+")
	return URL + query
}
