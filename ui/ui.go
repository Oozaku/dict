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

	"github.com/Oozaku/dict/dictapi"
)

var reader bufio.Reader

// Create reader that will take input from stdin
func init() {
	reader = *bufio.NewReader(os.Stdin)
}

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

// Print result
func PrintResults(entries []dictapi.Entry) {

	// Clear terminal first
	clearTerminal()

	// Each definition are listed with a increasing number
	counter := 1

	// For each result from search: print its 'name' and its phonetic
	for _, entry := range entries {
		fmt.Printf("%s\n", Bold(strings.ToUpper(entry.Word)))
		fmt.Printf("%s\n", Bold(entry.Phonetic).Italic())

		// For each meaning: print its function in the language
		for _, meaning := range entry.Meanings {
			fmt.Println(Italic(meaning.PartOfSpeech))

			// For each definition: list it with a increasing number and the definition
			for _, definition := range meaning.Definitions {
        list := fmt.Sprintf("%d.", counter)
				fmt.Printf("\t%s %s\n", Bold(list), definition.Def)
				counter++
			}

      // Separate each semantic group by a new line
      fmt.Println()
		}
	}
}

// Print welcome message
func PrintWelcome() {
	fmt.Println(Bold("Welcome to Dictionary API!!!"))
}
