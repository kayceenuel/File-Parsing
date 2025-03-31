package binary

import (
	"fmt"
	"io"
	"io/ioutil"
)

// Parser interface
type PlayerRecord struct {
	Name      string
	HighScore int
}

type parser struct{}

// parse reads binary data from the provided reader and returns a slice of playersRecord
func (p *parser) Parse(r io.Reader) ([]PlayersRecord, error) {
	// Read all data from the reader into a byte slice.
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read data: %v", err)
	}
	// ensure there's enough data for the endianess maker
	if len(data) < 2 {
		return nil, fmt.Errorf("insufficient data for endianness marker")
	}
}
