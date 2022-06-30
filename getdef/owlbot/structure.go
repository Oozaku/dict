package owlbot

// Word has name, phonetic and a list of meanings
type Word struct {
	Name     string    `json:"word"`
	Phonetic string    `json:"pronunciation"`
	Meanings []Meaning `json:"definitions"`
}

// Meaning has definition, emoji, example of usage, image and a part of speech
type Meaning struct {
	Definition   string `json:"definition"`
	Emoji        string `json:"emoji"`
	Example      string `json:"example"`
	ImageUrl     string `json:"image_url"`
	PartOfSpeech string `json:"type"`
}
