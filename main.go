package main

import (
	"flag"
	"strings"

	"github.com/dnozdrin/piglatin/text"
	_ "github.com/dnozdrin/piglatin/text/parser/regex"
	_ "github.com/dnozdrin/piglatin/text/translation/english"
)

func main() {
	var lang, source, target string

	flag.StringVar(&source, "source", "", "The source file path.")
	flag.StringVar(&target, "target", "", "The target file path.")
	flag.StringVar(&lang, "lang", "en", "The source language key.")
	flag.Parse()

	main := text.NewMainService(
		strings.TrimSpace(source),
		strings.TrimSpace(target),
		text.NewProcessor(lang),
	)
	main.Run()
}
