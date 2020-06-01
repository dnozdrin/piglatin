package text

import (
	"log"
	"strings"

	"github.com/dnozdrin/piglatin/text/parser"
	"github.com/dnozdrin/piglatin/text/translation"
)

// Processor represents a wrapper for text parsing and translation
type LineProcessor struct {
	translator translation.Translator
	parser     parser.Parser
}

// NewProcessor will return a pointer to a new Processor
func NewProcessor(lang string) *LineProcessor {
	psr, err := parser.NewParser(lang)
	if err != nil {
		log.Fatal(err)
	}
	tr, err := translation.NewTranslator(lang)
	if err != nil {
		log.Fatal(err)
	}

	return &LineProcessor{
		translator: tr,
		parser:     psr,
	}
}

type token struct {
	index int
	text  string
}

// Process will translate the given text
func (tp *LineProcessor) Process(line string) string {
	tokens := tp.parser.Parse(line)
	length := len(tokens)
	translated := make([]string, length, length)

	ch := make(chan token)
	defer close(ch)

	for i, val := range tokens {
		go func(i int, val string) {
			ch <- token{i, tp.translator.Translate(val)}
		}(i, val)
	}

	for i := 0; i < length; i++ {
		result := <-ch
		translated[result.index] = result.text
	}

	return strings.Join(translated, "")
}
