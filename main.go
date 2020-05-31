package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	lang "github.com/dnozdrin/piglatin/language"
	"github.com/dnozdrin/piglatin/parser"

	_ "github.com/dnozdrin/piglatin/language/english"
	_ "github.com/dnozdrin/piglatin/parser/regex"
)

// todo:
// - pass all linters
// todo: consider multithread
// todo: benchmarks
// todo: big file handling
// todo: add parser priority
// todo: check type pointer / value in code and tests
func main() {
	var result string
	langKey := *flag.String("lang", "en", "The source language key.")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	psr, err := parser.NewParser(langKey)
	if err != nil {
		log.Printf("reading input: %s", err)
		os.Exit(1)
	}

	tr, err := lang.NewTranslator(langKey)
	if err != nil {
		log.Printf("translation: %s", err)
		os.Exit(1)
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
		log.Printf("reading input: %s", err)
		os.Exit(1)
	}
}
