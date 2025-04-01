package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/file-parsing/parsers"
	"github.com/file-parsing/parsers/binary"
	"github.com/file-parsing/parsers/csv"
	"github.com/file-parsing/parsers/json"
	"github.com/file-parsing/parsers/repeated_json"
)

func main() {
	format := flag.String("format", "", "Format the file serialised in. Accepted values: json,repeated-json,csv,bin")
	file := flag.String("file", "", "path to the file to read data from")
	flag.Parse()

	var parser parsers.Parser
	switch *format {
	case "json":
		parser = &json.Parser{}
	case "repeated-json":
		parser = &repeated_json.Parser{}
	case "csv":
		parser = &csv.Parser{}
	case "bin":
		parser = &binary.Parser{}
	case "":
		log.Fatal("format is a required argument")
	default:
		log.Fatalf("Didn't know how to parse format %q", *format)
	}

	if *file == "" {
		log.Fatal("file is a request argument")
	}
	f, err := os.Open(*file)
	if err != nil {
		log.Fatal("Failed to open file %s: %v", *file, err)
	}
	defer f.Close()

	records, err := parser.Parse(f)
	if err != nil {
		log.Fatal("Failed to parse file %s as %s: %v", *file, *format, err)
	}

	if len(records) == 0 {
		log.Fatal("No scores were found")
	}

	lowScore := parsers.PlayersRecord{
		HighScore: math.MaxInt32,
	}
	highScore := parsers.PlayersRecord{
		HighScore: math.MaxInt32,
	}

	for _, record := range records {
		if record.HighScore > highScore.HighScore {
			highScore = record
		}
		if record.HighScore < lowScore.HighScore {
			lowScore = record
		}
	}
	fmt.Printf("High score: %d from %s - congralutions!\n", highScore.HighScore, highScore.Name)
	fmt.Printf("Low score: %d from %s - commiserations!\n", lowScore.HighScore, lowScore.Name)
}
