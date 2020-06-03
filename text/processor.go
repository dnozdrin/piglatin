package text

import (
	"strings"

	"github.com/dnozdrin/piglatin/text/parser"
	"github.com/dnozdrin/piglatin/text/translation"
)

// LineProcessor represents a wrapper for text parsing and translation
type LineProcessor struct {
	translator translation.Translator
	parser     parser.Parser
}

// NewProcessor will return a pointer to a new Processor
func NewProcessor(tr translation.Translator, psr parser.Parser) *LineProcessor {
	return &LineProcessor{
		translator: tr,
		parser:     psr,
	}
}

// Process will translate the given text
func (proc *LineProcessor) Process(line string) string {
	tokens := proc.parser.Parse(line)
	length := len(tokens)
	translated := make([]string, length, length)

	for i, val := range tokens {
		translated[i] = proc.translator.Translate(val)
	}

	return strings.Join(translated, "")
}
