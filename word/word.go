/*
Package word contains the struct that is used to hold the properties a word and
all modules that handles different providers have to convert to this struct.

The properties that a word holds are:
  - Its ortography ("Name")
  - Phonetic
  - Meaning
*/
package word

// Word holds the phonetics and its meaning
type Word struct {
	// Name is the ortography of a word
	Name string
	// Phonetics of the word
	Phonetics []Phonetic
	// Meanings holds the definitions of the word grouped by their part of speech
	Meanings map[string][]Meaning
}
