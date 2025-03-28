package binary

import (
	
)

type parser struct{}

func (p *parser) Parse(r io.Reader) ([]parsers.PlayersRecord, error) {
	bufRead := bufio.NewReader(r)

	var records := parsers.playersRecord
}