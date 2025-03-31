package binary

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
)

// Parser interface (assumed from a parsers package)
type PlayersRecord struct {
	Name      string
	HighScore int
}

type parser struct{}

// Parse reads binary data from the provided reader and returns a slice of PlayersRecord.
func (p *parser) Parse(r io.Reader) ([]PlayersRecord, error) {
	// Read all data from the reader into a byte slice
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read data: %v", err)
	}

	// Ensure there’s enough data for the endianness marker
	if len(data) < 2 {
		return nil, fmt.Errorf("insufficient data for endianness marker")
	}

	// Determine endianness based on the first two bytes
	var endian binary.ByteOrder
	switch {
	case bytes.Equal(data[:2], []byte{0xFE, 0xFF}):
		endian = binary.BigEndian
	case bytes.Equal(data[:2], []byte{0xFF, 0xFE}):
		endian = binary.LittleEndian
	default:
		return nil, fmt.Errorf("invalid endianness marker")
	}

	// Skip the endianness marker
	data = data[2:]

	// Slice to store parsed records
	var records []PlayersRecord

	// Process records until no data remains
	for len(data) > 0 {
		// Ensure there’s enough data for the score (4 bytes)
		if len(data) < 4 {
			return nil, fmt.Errorf("insufficient data for score")
		}

		// Read and interpret the score based on endianness
		score := int32(endian.Uint32(data[:4]))
		data = data[4:]

		// Find the null terminator for the name
		nullIndex := bytes.IndexByte(data, 0)
		if nullIndex == -1 {
			return nil, fmt.Errorf("missing null terminator for name")
		}

		// Extract and decode the name
		nameBytes := data[:nullIndex]
		name := string(nameBytes) // UTF-8 decoding is implicit in Go’s string conversion
		data = data[nullIndex+1:]

		// Add the record to the slice
		records = append(records, PlayersRecord{
			Name:      name,
			HighScore: int(score),
		})
	}

	return records, nil
}
