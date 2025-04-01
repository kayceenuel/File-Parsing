package json

import (
	"encoding/json"
	"io"

	"github.com/file-parsing/parsers"
)

type Parser struct{}

func (p *Parser) Parse(r io.Reader) ([]parsers.PlayersRecord, error) {
	var records []parsers.PlayersRecord
	if err := json.NewDecoder(r).Decode(&records); err != nil {
		return nil, err
	}
	return records, nil
}
