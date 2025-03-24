package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"

	"github.com/file-parsing/parsers"
)

type parser struct{}

func (p *parser) Parse(r io.Reader) ([]parsers.PlayersRecord, error) {
	// Create a new CSV reader
	csvReader := csv.NewReader(r)

	// Read the CSV headers
	headerRecord, err := csvReader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read headers: %v", err)
	}

	// Find indices of "name" and "high score" headers
	nameIndex := -1
	highScoreIndex := -1
	for i, col := range headerRecord {
		if col == "name" {
			nameIndex = i
		} else if col == "high score" {
			highScoreIndex = i
		} else {
			return nil, fmt.Errorf("unexpected header %q - expected %q and %q", col, "name", "high score")
		}
	}
	if nameIndex == -1 || highScoreIndex == -1 {
		return nil, fmt.Errorf("incorrect headers - expected to find %q and %q", "name", "high score")
	}

	// Parse records into a slice of PlayersRecord
	var records []parsers.PlayersRecord
	for {
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break // End of file reached, exit the loop
			}
			return nil, fmt.Errorf("failed to read record: %v", err)
		}

		// Ensure the record has enough columns
		if len(record) <= nameIndex || len(record) <= highScoreIndex {
			return nil, fmt.Errorf("record has insufficient columns")
		}

		// Extract and convert fields
		name := record[nameIndex]
		highScoreStr := record[highScoreIndex]
		highScore, err := strconv.Atoi(highScoreStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse high score %q: %v", highScoreStr, err)
		}

		// Append the parsed record to the slice
		records = append(records, parsers.PlayersRecord{
			Name:      name,
			HighScore: highScore,
		})
	}

	return records, nil
}
