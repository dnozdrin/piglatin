package text

import (
	"strings"
	"testing"
	"time"
)

type translatorMock struct{}

func (mock *translatorMock) Translate(word string) string {
	time.Sleep(20 * time.Millisecond)

	return word + "!"
}

type parserMock struct{}

func (p *parserMock) Parse(text string) []string {
	return strings.Fields(text)
}

var canHandleMock = func(key string) bool {
	return true
}

func (p *parserMock) CanHandle(key string) bool {
	return canHandleMock(key)
}

func BenchmarkHandleTokens(b *testing.B) {
	processorMock := NewProcessor(&translatorMock{}, &parserMock{})
	for i := 0; i < b.N; i++ {
		processorMock.Process("quiet quiet quiet")
	}
}

var testLines = []struct {
	in, out string
}{
	{"latin test 1", "latin!test!1!"},
	{"are test 2", "are!test!2!"},
	{"quiet quiet quiet", "quiet!quiet!quiet!"},
}

func TestProcess(t *testing.T) {
	processorMock := NewProcessor(&translatorMock{}, &parserMock{})
	for _, tt := range testLines {
		t.Run(tt.in, func(t *testing.T) {
			result := processorMock.Process(tt.in)
			if result != tt.out {
				t.Errorf("got %q, want %q", result, tt.out)
			}
		})
	}
}
