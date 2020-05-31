package regex

import (
	"reflect"
	"regexp"
	"testing"
)

var languages = []struct {
	in  string
	out bool
}{
	{"en", true},
	{"ua", true},
	{"ru", true},
	{"fr", false},
}

func TestCanHandle(t *testing.T) {
	expr, _ := regexp.Compile(``)
	regex := NewRegex(expr)
	for _, tt := range languages {
		t.Run(tt.in, func(t *testing.T) {
			s := regex.CanHandle(tt.in)
			if s != tt.out {
				t.Errorf("got %t, want %t", s, tt.out)
			}
		})
	}
}

var samples = []struct {
	in  string
	out []string
}{
	{"The quick brown fox jumps over a lazy dog.",
		[]string{"The", " ", "quick", " ", "brown", " ", "fox", " ", "jumps", " ", "over", " ", "a", " ", "lazy", " ", "dog", "."},
	},
	{"latin-banana", []string{"latin", "-", "banana"}},
	{"По-моему, это неплохой тест", []string{"По", "-", "моему", ", ", "это", " ", "неплохой", " ", "тест"}},
	{"Україна — унітарна держава.", []string{"Україна", " — ", "унітарна", " ", "держава", "."}},
}

func TestParse(t *testing.T) {
	expr, _ := regexp.Compile(pattern)
	regex := NewRegex(expr)
	for _, tt := range samples {
		t.Run(tt.in, func(t *testing.T) {
			s := regex.Parse(tt.in)
			if !reflect.DeepEqual(s, tt.out) {
				t.Errorf("got %q, want %q", s, tt.out)
			}
		})
	}
}

func TestNewParser(t *testing.T) {
	t.Run("regexp property", func(t *testing.T) {
		expr := regexp.MustCompile(``)
		regex := NewRegex(expr)

		if regex.expr != expr {
			t.Errorf("got %q want %q given, %q", expr, expr, regex.expr)
		}
	})
}
