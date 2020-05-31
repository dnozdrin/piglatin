package regex

import (
	"github.com/dnozdrin/piglatin/parser"
	"regexp"
)

const pattern = `([^\s—'[:punct:]]+)+([\s—'[:punct:]]*)`

func init() {
	parser.Register(NewRegex(regexp.MustCompile(pattern)))
}

type Regex struct {
	expr     *regexp.Regexp
	supports map[string]struct{}
}

// NewRegex returns a pointer to the regex parser
func NewRegex(expr *regexp.Regexp) *Regex {
	return &Regex{
		expr: expr,
		supports: map[string]struct{}{
			"en": {},
			"ua": {},
			"ru": {},
		},
	}
}

// CanHandle states if the parser can handler a language
// with the given key
func (r *Regex) CanHandle(key string) bool {
	_, ok := r.supports[key]
	return ok
}

// Parse will parse the given string and return all matches
// in a slice
func (r *Regex) Parse(text string) []string {
	matches := r.expr.FindAllStringSubmatch(text, -1)
	result := make([]string, 0)
	for _, m := range matches {
		if len(m) == 3 {
			result = append(result, m[1])
			if len(m[2]) > 0 {
				result = append(result, m[2])
			}
		}
	}

	return result
}
