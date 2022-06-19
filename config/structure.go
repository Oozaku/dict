package config

type Config struct {
	// AnkiCsvLocation holds the location to csv file that is used to import
	// inside Anki
	AnkiCsvLocation string `json:"anki-csv-location"`
	// AnkiMediaFolder holds the folder location where Anki saves all media files
	// (images and sounds).
	AnkiMediaFolder string `json:"anki-media-folder"`
}
