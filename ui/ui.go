package ui

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	. "github.com/logrusorgru/aurora"

	"github.com/Oozaku/dict/word"
)

var reader bufio.Reader = *bufio.NewReader(os.Stdin)

// Clear terminal screen
func clearTerminal() {
	switch runtime.GOOS {
	case "darwin", "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// Removes empty strings
func filterNonEmpty(words []string) []string {
	final := make([]string, 0)
	for _, word := range words {
		if len(word) > 0 {
			final = append(final, word)
		}
	}
	return final
}

// Get input from user and return a slice of words
func GetEntryFromUser() []string {
	fmt.Printf("Search for: ")
	word, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	word = strings.TrimSpace(word)
	listWords := strings.Split(word, " ")
	return filterNonEmpty(listWords)
}

// Print all meanings grouped by part of speech
func printMeanings(word word.Word) {

	// For each meaning: print its part of speech and its definition
  for _, meaning := range word.Meanings {
    fmt.Printf("(%s) %s\n", meaning.PartOfSpeech, meaning.Definition)
  }
}

// Print all non empty phonetics
func printPhonetics(phonetics []word.Phonetic) {
	for _, phonetic := range phonetics {
		if len(phonetic.Text) != 0 {
			fmt.Printf("%s\n", Bold(phonetic.Text).Italic())
		}
	}
}

// Print result
func PrintResults(words []word.Word) {

	// Clear terminal first
	clearTerminal()

	// For each result from search: print its 'name' and its phonetic
	for _, word := range words {
		fmt.Printf("%s\n", Bold(strings.ToUpper(word.Name)))

		// Print all phonetics
		printPhonetics(word.Phonetics)

		// Print all meanings grouped by part of speech
		printMeanings(word)
	}
}

// Print welcome message
func PrintWelcome() {
	fmt.Println(Bold("Welcome to Dict"))
}
