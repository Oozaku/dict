/*
Package getdef holds the implementation that searches and abstracts
the data retrieved to one that the main module can digest.
*/
package getdef

import (
	"github.com/Oozaku/dict/config"
	"github.com/Oozaku/dict/getdef/dictionaryapi"
	"github.com/Oozaku/dict/word"
)

// Interface specifies that all api clients must have method SearchDefinition
// that takes a list of strings and return a list of words
type Client interface {
  SearchDefinition(words []string) ([]word.Word, error)
}

func RetrieveMapOfClients(configuration config.Config) map[string]Client {
  return map[string]Client{
    "dictionaryapi": dictionaryapi.DictionaryApi{},
    "owlbot": configuration.Owlbot,
  }
}
