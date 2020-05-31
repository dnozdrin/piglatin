package logger

import (
	"log"
	"os"
	"path/filepath"
)

// Error logger
var Error *log.Logger

func init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Print(err)
	}

	absPath, err := filepath.Abs(dir + "/log")
	if err != nil {
		log.Printf("Error reading given path: %s", err)
	}

	file, err := os.OpenFile(absPath+"/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("Error opening file: %s", err)
	}

	Error = log.New(file, "\t", log.Ldate|log.Ltime|log.Lshortfile)
}
