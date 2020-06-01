package text

import (
	"testing"
	"time"
)

type translatorMock struct{}

func (mock *translatorMock) Translate(word string) string {
	time.Sleep(20 * time.Millisecond)

	return word
}

type parserMock struct{}

var parserFuncMock = func(text string) []string {
	return []string{}
}

func (p *parserMock) Parse(text string) []string {
	return parserFuncMock(text)
}

var canHandleMock = func(key string) bool {
	return true
}

func (p *parserMock) CanHandle(key string) bool {
	return canHandleMock(key)
}

func BenchmarkHandleTokens(b *testing.B) {
	processorMock := &LineProcessor{
		&translatorMock{},
		&parserMock{},
	}

	for i := 0; i < b.N; i++ {
		processorMock.Process("golang")
	}
}
