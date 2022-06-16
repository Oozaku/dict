/*
Package media implements the functionalities to download the audio.
*/
package media

import (
	"io/ioutil"
	"log"

	"github.com/Oozaku/dict/requests"
)

// DownloadAudio downloads the audio from url and saves it in path with the
// following permissions:
// - read and write to user and his group
// - read only to others
func DownloadAudio(url, path string) {

	// Downloads audio's binary
	audio, err := requests.Get(url)
	failInError(err)

	// Saves the file with read with correct permissions
	err = ioutil.WriteFile(path, audio, 664)
	failInError(err)
}

// failInError does nothing if there is no error, but it fails the program if
// there was an error
func failInError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
