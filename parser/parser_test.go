package parser

import "testing"

type testParser struct{}

var canHandleMock = func(key string) bool {
	return true
}

func (p *testParser) CanHandle(key string) bool {
	return canHandleMock(key)
}

var parseMock func(text string) []string

func (p *testParser) Parse(text string) []string {
	return parseMock(text)
}

func TestRegister(t *testing.T) {
	t.Run("parser registration", func(t *testing.T) {
		parsers = []Parser{}
		parserMock := &testParser{}
		Register(parserMock)

		if parsers == nil {
			t.Errorf("got %v, want []Parser", nil)
		}

		for _, parser := range parsers {
			if parser != parserMock {
				t.Errorf("got %t, want %t", parser, parserMock)
			}
		}
	})
	t.Run("parser registration panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("the code did not panic on nil registration")
			}
		}()

		Register(nil)
	})
}

func TestNewParser(t *testing.T) {
	t.Run("new parser", func(t *testing.T) {
		parserMock := &testParser{}
		parsers = []Parser{parserMock}

		parser, err := NewParser("dummy")
		if err != nil {
			t.Fatal("can not instantiate parser")
		}

		if parser != parserMock {
			t.Errorf("got %t, want %t", parser, parserMock)
		}
	})

	t.Run("test error on no parsers", func(t *testing.T) {
		parsers = []Parser{}
		parser, err := NewParser("dummy")
		if err == nil {
			t.Errorf("got nil, want error")
		}
		if parser != nil {
			t.Errorf("got %t, want nil", parser)
		}
	})
}
