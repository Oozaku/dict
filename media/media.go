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
func DownloadAudio(url, path string) error {

	// Downloads audio's binary
	headers := make(map[string]string)
	audio, err := requests.Get(url, headers)
	if err != nil {
		log.Println(err)
		return err
	}

	// Saves the file with read with correct permissions
	err = ioutil.WriteFile(path, audio, 0664)
	if err != nil {
		log.Println(err)
		return err
	}

	// No errors occurred
	return nil
}
