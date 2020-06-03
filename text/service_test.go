package text

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

type testInterface struct{}

func (ti *testInterface) Process(text string) string {
	return text + "111"
}

var samples = []struct {
	in, out string
}{
	{"latin", "latin111"},
	{"are", "are111"},
	{"quiet quiet quiet", "quiet quiet quiet111"},
}

func TestStdRun(t *testing.T) {
	for _, tt := range samples {
		t.Run(tt.in, func(t *testing.T) {
			service := NewMainService("", "", &testInterface{})
			service.reader = fakeInput(t, tt.in)

			r, w, err := os.Pipe()
			service.writer = w
			if err != nil {
				t.Fatal(err)
			}
			service.Run()
			w.Close()
			var buf bytes.Buffer
			io.Copy(&buf, r)

			result := strings.TrimSuffix(buf.String(), "\n")
			if result != tt.out {
				t.Errorf("got %q, want %q", result, tt.out)
			}
		})
	}
}

func fakeInput(t *testing.T, text string) *os.File {
	content := []byte(text)
	tmpFile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := tmpFile.Write(content); err != nil {
		t.Fatal(err)
	}

	if _, err := tmpFile.Seek(0, 0); err != nil {
		t.Fatal(err)
	}

	return tmpFile
}

func TestFilesRun(t *testing.T) {
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
			service := NewMainService(inputFile.Name(), outputFile.Name(), &testInterface{})
			err = service.Run()
			if err != nil {
				t.Errorf("got error on service run")
			}

			output, err := ioutil.ReadFile(outputFile.Name())
			if err != nil {
				t.Fatalf("unable to read file: %v", err)
			}

			result := string(output)
			result = strings.TrimSuffix(result, "\n")
			if result != tt.out {
				t.Errorf("got %q, want %q", result, tt.out)
			}
		})
	}
}

func TestFilesRunError(t *testing.T) {
	t.Run("test source error", func(t *testing.T) {
		service := NewMainService("dummy1", "", &testInterface{})
		err := service.Run()

		if err == nil {
			t.Error("got nil, want error")
		}
	})
}
