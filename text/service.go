package text

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type MainService struct {
	source, target string
	reader         io.ReadCloser
	writer         io.WriteCloser
	processor      Processor
}

// Processor represents a wrapper for text parsing and translation
type Processor interface {
	Process(string) string
}

// NewMainService constructor
func NewMainService(source, target string, processor Processor) *MainService {
	return &MainService{
		source:    source,
		target:    target,
		reader:    os.Stdin,
		writer:    os.Stdout,
		processor: processor,
	}
}

// Run handles the main login of the app
func (ms *MainService) Run() {
	if ms.source != "" {
		path, err := filepath.Abs(ms.source)
		if err != nil {
			log.Print(err)
		}

		ms.reader, err = os.Open(path)
		if err != nil {
			log.Panicf("An error occured while opening file: %s", err)
		}
		defer ms.reader.Close()
	}

	if ms.target != "" {
		path, err := filepath.Abs(ms.target)
		if err != nil {
			log.Print(err)
		}
		ms.writer, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			log.Panicf("An error occured while opening file: %s", err)
		}

		defer ms.writer.Close()
	}

	scanner := bufio.NewScanner(ms.reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if line == ":q" {
			break
		}
		line = ms.processor.Process(line)
		_, err := fmt.Fprintln(ms.writer, line)
		if err != nil {
			log.Println(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
