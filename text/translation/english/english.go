package english

import (
	"regexp"
	"strings"
	"unicode"

	lang "github.com/dnozdrin/piglatin/text/translation"
)

const (
	key         = "en"
	wordPattern = `^[A-Za-z']+$`
)

var expr *regexp.Regexp

func init() {
	lang.Register(key, NewEnglish())
	expr = regexp.MustCompile(wordPattern)
}

// NewEnglish will return a pointer to English struct
func NewEnglish() *English {
	return &English{}
}

// English represents a translator to Pig Latin from English language
type English struct{}

// Translate will translate a English word to Pig Latin.
// As far as English alphabet is completely in the ASCII Table, it
// works with bytes.
// The next rules are considered:
// - Ensures proper capitalization
// - Correct upper case and lower case formatting
// - Correctly translates "qu" (e.g., ietquay instead of uietqay)
// - Differentiates between "Y" as vowel and "Y" as consonant
// (e.g. yellow = ellowyay and style = ylestay)
// - Correctly translates contractions
// - Words may consist of alphabetic characters only (A-Z and a-z)
func (eng *English) Translate(word string) string {
	if !isTranslatable(word) {
		return word
	}
	if isVowel(word[0]) {
		return word + "yay"
	}

	result := word

	isFirstUpper := unicode.IsUpper(rune(word[0]))
	if isFirstUpper {
		word = strings.ToLower(string(word[0])) + word[1:]
	}

	for i, letter := range []byte(word[1:]) {
		if isVowel(letter) {
			if isEqual(letter, "U") && i >= 0 && isEqual(word[i], "Q") {
				result = word[i+2:] + word[0:i+2] + "ay"
				break
			}

			result = word[i+1:] + word[0:i+1] + "ay"
			break
		}
		if isEqual(letter, "Y") {
			result = word[i+1:] + word[0:i+1] + "ay"
			break
		}
	}

	if isFirstUpper {
		result = strings.Title(result)
	}

	return result
}

func isTranslatable(word string) bool {
	return expr.MatchString(word)
}

func isVowel(l byte) bool {
	_, ok := vowels[l]
	return ok
}

var vowels = map[byte]struct{}{
	'a': {},
	'e': {},
	'i': {},
	'o': {},
	'u': {},
	'A': {},
	'E': {},
	'I': {},
	'O': {},
	'U': {},
}

func isEqual(letter byte, check string) bool {
	return strings.ToLower(string(letter)) == strings.ToLower(check)
}
