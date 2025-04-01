package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/file-parsing/parsers"
	"github.com/file-parsing/parsers/binary"
	"github.com/file-parsing/parsers/csv"
	"github.com/file-parsing/parsers/json"
	"github.com/file-parsing/parsers/repeated_json"
)

func main() {
	// Define command-line flags with improved descriptions
	format := flag.String("format", "", "Format of the file. Accepted values: json, repeated-json, csv, binary")
	file := flag.String("file", "", "Path to the file to read data from")
	flag.Parse()

	// Select parser based on format
	var parser parsers.Parser
	switch *format {
	case "json":
		parser = &json.Parser{}
	case "repeated-json":
		parser = &repeated_json.Parser{}
	case "csv":
		parser = &csv.Parser{}
	case "binary":
		parser = &binary.Parser{}
	case "":
		log.Fatal("format is a required argument")
	default:
		log.Fatalf("Didn't know how to parse format %q", *format)
	}

	// Check if file is provided
	if *file == "" {
		log.Fatal("file is a required argument")
	}

	// Open the file
	f, err := os.Open(*file)
	if err != nil {
		log.Fatalf("Failed to open file %s: %v", *file, err) // Use log.Fatalf for formatting
	}
	defer f.Close()

	// Parse the file
	records, err := parser.Parse(f)
	if err != nil {
		log.Fatalf("Failed to parse file %s as %s: %v", *file, *format, err) // Use log.Fatalf
	}

	// Check if records are empty
	if len(records) == 0 {
		log.Fatal("No scores were found")
	}

	// Initialize highScore and lowScore with the first record
	highScore := records[0]
	lowScore := records[0]

	// Loop through remaining records to find highest and lowest scores
	for _, record := range records[1:] {
		if record.HighScore > highScore.HighScore {
			highScore = record
		}
		if record.HighScore < lowScore.HighScore {
			lowScore = record
		}
	}

	// Print results with corrected spelling
	fmt.Printf("High score: %d from %s - congratulations!\n", highScore.HighScore, highScore.Name)
	fmt.Printf("Low score: %d from %s - commiserations!\n", lowScore.HighScore, lowScore.Name)
}
