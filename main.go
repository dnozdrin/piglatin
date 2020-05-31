package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/dnozdrin/piglatin/logger"

	lang "github.com/dnozdrin/piglatin/language"
	"github.com/dnozdrin/piglatin/parser"

	_ "github.com/dnozdrin/piglatin/language/english"
	_ "github.com/dnozdrin/piglatin/parser/regex"
)

func main() {
	var result, langKey string

	flag.StringVar(&langKey, "lang", "en", "The source language key.")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	psr, err := parser.NewParser(langKey)
	if err != nil {
		logger.Error.Fatalf("reading input: %s", err)
	}

	tr, err := lang.NewTranslator(langKey)
	if err != nil {
		logger.Error.Fatal(err)
	}
	for scanner.Scan() {
		result = ""
		line := scanner.Text()
		if line == ":q" {
			break
		}
		words := psr.Parse(line)
		for _, word := range words {
			result += tr.Translate(word)
		}
		fmt.Println(result)
	}
	if err := scanner.Err(); err != nil {
		logger.Error.Fatalf("reading input: %s", err)
	}
}
