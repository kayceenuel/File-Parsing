package binary

import (
	"io"
)

// Parser interface
type PlayerRecord struct {
	Name      string
	HighScore int
}

type parser struct{}

func (p *parser) Parse(r io.Reader) ([]PlayersRecord, error) {

}
