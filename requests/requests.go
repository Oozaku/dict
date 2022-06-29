/*
Package requests have functions to make external requests
*/
package requests

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/Oozaku/dict/errors"
)

// Get makes a GET request to url using custom headers and returns a slice of
// bytes and an error.
func Get(url string, headers map[string]string) ([]byte, error) {

	// Create client
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Add custom headers
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	// Get data from API
	resp, err := client.Do(req)
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
		return []byte{}, &errors.ReqError{Message: string(body),
			Code: resp.StatusCode}
	}

	return body, nil
}
