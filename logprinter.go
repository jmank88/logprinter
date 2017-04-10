package logprinter

import (
	"fmt"
	"github.com/go-kit/kit/log"
)

type Printer struct {
	logger log.Logger
	key    string
}

// NewPrinter returns a new Printer which logs messages as values of key.
func NewPrinter(logger log.Logger, key string) *Printer {
	return &Printer{
		logger: logger,
		key:    key,
	}
}

func (p *Printer) Print(s string) error {
	return p.logger.Log(p.key, s)
}

func (p *Printer) Printf(s string, a ...interface{}) error {
	return p.logger.Log(p.key, fmt.Sprintf(s, a...))
}
