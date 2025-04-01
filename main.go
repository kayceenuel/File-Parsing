package main

import (
	"flag"
	"log"

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
}
