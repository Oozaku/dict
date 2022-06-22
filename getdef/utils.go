package getdef

import "github.com/Oozaku/dict/word"

// JoinWords join words with same name
func JoinWords(words []word.Word) []word.Word {

	// Map to check if new word is already present
	mapping := make(map[string]word.Word)

	for _, word := range words {

		if w, ok := mapping[word.Name]; ok {
			// word is already tracked: append its meanings into word in map
			for _, meaning := range word.Meanings {
				w.Meanings = append(w.Meanings, meaning)
			}
			mapping[word.Name] = w

		} else {
			// word not tracked: track it
			mapping[word.Name] = word
		}
	}

  // Convert map into list
	var result []word.Word
	for _, word := range mapping {
		result = append(result, word)
	}

  // Return list
	return result
}
