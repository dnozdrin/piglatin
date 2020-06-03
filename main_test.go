package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var samples = []struct {
	in, out string
}{
	{"latin", "atinlay"},
	{"are", "areyay"},
	{
		"Lorem ipsum dolor sit amet.",
		"Oremlay ipsumyay olorday itsay ametyay.",
	},
}

func TestMainFromFiles(t *testing.T) {
	for _, tt := range samples {
		t.Run(tt.in, func(t *testing.T) {
			inputFile, err := ioutil.TempFile("", "prefix")
			if err != nil {
				t.Fatal(err)
			}
			outputFile, err := ioutil.TempFile("", "prefix")
			if err != nil {
				t.Fatal(err)
			}

			defer os.Remove(inputFile.Name())
			defer os.Remove(outputFile.Name())

			err = ioutil.WriteFile(inputFile.Name(), []byte(tt.in), 0755)
			if err != nil {
				t.Fatalf("unable to write file: %v", err)
			}

			os.Args = []string{"piglatin",
				fmt.Sprintf("-source=%s", inputFile.Name()),
				fmt.Sprintf("-target=%s", outputFile.Name()),
			}
			main()

			output, err := ioutil.ReadFile(outputFile.Name())
			result := string(output)
			result = strings.TrimSuffix(result, "\n")
			if err != nil {
				t.Fatalf("unable to read file: %v", err)
			}
			if result != tt.out {
				t.Errorf("got %q, want %q", result, tt.out)
			}
		})
	}
}

func TestMainErrorLogging(t *testing.T) {
	const errorsExpected = 2

	t.Run("test errors on wrong lang", func(t *testing.T) {
		origLogFatal := logFatal
		defer func() { logFatal = origLogFatal }()

		errors := make([]interface{}, 0)
		logFatal = func(v ...interface{}) {
			errors = append(errors, v)
		}

		os.Args = []string{"piglatin",
			fmt.Sprintf("-lang=%s", "wrong"),
		}
		main()

		errorsActual := len(errors)
		if errorsActual != errorsExpected {
			t.Errorf("excepted %d errors, got %d", errorsExpected, errorsActual)
		}
	})
}

func TestMainErrorOnWrongArgument(t *testing.T) {
	t.Run("test errors on wrong argument", func(t *testing.T) {
		origLogFatal := logFatal
		defer func() { logFatal = origLogFatal }()

		var errReturned bool
		logFatal = func(v ...interface{}) {
			errReturned = true
		}

		os.Args = []string{"piglatin", "dummy"}
		main()

		if !errReturned {
			t.Error("excepted error on wrong argument")
		}
	})
}
