package language

import "fmt"

var translators = make(map[string]Translator)

// Translator represents an interface for translators to Pig Latin
// from different languages
type Translator interface {
	Translate(text string) string
}

// Register will register a translator with the provided
// language key
func Register(key string, t Translator) {
	if t == nil {
		panic("piglatin: Register translator is nil")
	}
	if _, dup := translators[key]; dup {
		panic("piglatin: Register called twice for translator " + key)
	}
	translators[key] = t
}

// NewTranslator will return a translator registered by the provided
// key or an error if a translator with such key is not registered
func NewTranslator(key string) (Translator, error) {
	t, ok := translators[key]
	if !ok {
		return nil, fmt.Errorf("piglatin: translator with key %s is not registered", key)
	}

	return t, nil
}
