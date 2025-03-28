package binary

import (
	"bufio"
	"io"

	"github.com/file-parsing/parsers"
)

type parser struct{}

func (p *parser) Parse(r io.Reader) ([]parsers.PlayersRecord, error) {
	bufRead := bufio.NewReader(r)

	var records []parsers.PlayersRecord
}
