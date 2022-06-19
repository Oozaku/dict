# Dictionary

## About

This project is a dictionary/language learner tool that will let you search
for meaning/translation of words and it will save the results inside a file,
making it easy to import inside a ANKI deck.

## Dependencies

- [go](https://go.dev/doc/install)

## Install

Make sure that GOPATH is in your path and run the following command.

~~~bash
go install github.com/Oozaku/dict@latest
~~~

If you don't have the GOPATH, run `go env | grep GOPATH` and copy and paste the
value inside double quotes to your PATH.

## Before Usage

To use this program, you need to set the configuration file at
`~/.config/dict/dict.json`. The code snippet bellow shows an example:

~~~json
{
  "anki-media-folder": "/home/oozaku/.local/share/Anki2/User 1/collection.media",
  "anki-csv-location": "/home/oozaku/Documents/anki/anki.csv"
}
~~~

The field `anki-media-folder` is where Anki saves all medias used in your decks
and `anki-csv-location` is the file location that you will use later to import
to Anki.

The Anki's media folder is a little trick to find, you have to find where
`collection.media` is located. On Linux, it probably is at
`~/.local/share/Anki2/<your_profile>/collection.media`.

## Usage

Run `dict` at your terminal and then you can search for expressions in the
configured API. Anki's import file and all audios are saved automatically inside
the configured locations.

To import the searched words into Anki, open Anki and go to File > Import...,
select your import file and set the first field as front and the second field as
back.

One important note is that when Anki finds two entries with same name, it
overwrites the old one with the new one. So it's important that you import
inside a new deck to avoid overwritting you existing cards.

## TO DO

- Features
  - [x] Integrate an API
  - [x] Save textual results in a supported format to Anki
  - [x] Download audios
  - [x] Save audio results inside file to be imported in Anki
  - [x] Define configuration file
  - [ ] Create interface to abstract each API client
  - [ ] Improve cli/tui
  - [ ] Lookup firstly at memory, if not found, search in API
  - [ ] Lookup in each API until find answer
- Documentation:
  - Improve README.md:
    - [ ] Add instructions on how to build/install
    - [ ] Add documentation about how it works
  - Code:
    - [ ] Add comments
