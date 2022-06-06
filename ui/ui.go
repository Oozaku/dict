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

	// Counter used to list each definition
	counter := 1

	// For each meaning: print its function in the language
	for partOfSpeech, meanings := range word.Meanings {
		fmt.Println(Italic(partOfSpeech))

		// For each definition: list it with a increasing number and the definition
		for _, meaning := range meanings {
			list := fmt.Sprintf("%d.", counter)
			fmt.Printf("\t%s %s\n", Bold(list), meaning.Definition)
			counter++
		}

		// Separate each semantic group by a new line
		fmt.Println()
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
