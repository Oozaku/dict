/*
Package getdef holds the implementation that searches and abstracts
the data retrieved to one that the main module can digest.
*/
package getdef

import (
	"github.com/Oozaku/dict/getdef/dictionaryapi"
	"github.com/Oozaku/dict/word"
)

// All clients should return a list of words followed by an error
type ProviderRetriever map[string]func([]string) ([]word.Word, error)

// Map with all clients
var GetProvider map[string]func([]string) ([]word.Word, error) = ProviderRetriever{
	"dictionaryapi": dictionaryapi.SearchDefinition,
}
