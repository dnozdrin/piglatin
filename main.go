package main

import (
	"errors"
	"flag"
	"log"
	"strings"

	"github.com/dnozdrin/piglatin/text/parser"
	"github.com/dnozdrin/piglatin/text/translation"

	"github.com/dnozdrin/piglatin/text"
	_ "github.com/dnozdrin/piglatin/text/parser/regex"
	_ "github.com/dnozdrin/piglatin/text/translation/english"
)

var lang, source, target string

func init() {
	regStringVar(&lang, "lang", "en", "The source language key.")
	regStringVar(&source, "source", "", "The source file path.")
	regStringVar(&target, "target", "", "The target file path.")
}

func initFlags() {
	lang = getStringFlag("lang")
	source = getStringFlag("source")
	target = getStringFlag("target")
}

func regStringVar(p *string, name string, value string, usage string) {
	if flag.Lookup(name) == nil {
		flag.StringVar(p, name, value, usage)
	}
}

func getStringFlag(name string) string {
	return flag.Lookup(name).Value.(flag.Getter).Get().(string)
}

var logFatal = log.Fatal

func main() {
	flag.Parse()
	initFlags()

	if flag.Arg(0) != "" {
		logFatal(errors.New("the app does not accept any arguments"))
	}

	psr, err := parser.NewParser(lang)
	if err != nil {
		logFatal(err)
	}
	tr, err := translation.NewTranslator(lang)
	if err != nil {
		logFatal(err)
	}

	textProcessor := text.NewProcessor(tr, psr)
	mainService := text.NewMainService(
		strings.TrimSpace(source),
		strings.TrimSpace(target),
		textProcessor,
	)
	if err := mainService.Run(); err != nil {
		log.Print(err)
	}
}
