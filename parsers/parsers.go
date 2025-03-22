package parsers

import "io"

// PlayersRecord holds a player's name and high score.
type PlayersRecord struct {
	Name      string `json:"name"`
	HighScore int    `json:"high_score"`
}

// Parser defines the interface for parsing player records from a data source.
type Parser interface {
	// Parse reads data from the provided reader and returns a slice of player records.
	Parse(r io.Reader) ([]PlayersRecord, error)
}
