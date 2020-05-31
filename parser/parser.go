package parser

import "fmt"

var parsers []Parser

type Parser interface {
	CanHandle(key string) bool
	Parse(text string) []string
}

// Register will register a new provider
func Register(p Parser) {
	if p == nil {
		panic("piglatin: Register parser is nil")
	}

	parsers = append(parsers, p)
}

// NewParser returns a parser that supports with a language
// represented by the given language key.
// Returns error in case no parsers support a language with
// the given key
func NewParser(key string) (Parser, error) {
	for _, parser := range parsers {
		if parser.CanHandle(key) {
			return parser, nil
		}
	}

	return nil, fmt.Errorf("piglatin: no parsers that supports lang key %s ", key)
}
