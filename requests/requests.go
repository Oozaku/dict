/*
Package requests have functions to make external requests
*/
package requests

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Oozaku/dict/errors"
)

// Make a get request to url and return the raw result
func Get(url string) ([]byte, error) {

	// Get data from API
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	// Close connection at the end
	defer resp.Body.Close()

	// Read all data from response and return it
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Status code is an error: return error
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return []byte{}, errors.NoResults{ErrorMessage: string(body), ErrorCode: 404}
	}

	return body, nil
}
