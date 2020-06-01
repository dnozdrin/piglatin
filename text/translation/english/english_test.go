package english

import (
	"reflect"
	"testing"
)

var samples = []struct {
	in, out string
}{
	{"latin", "atinlay"},
	{"are", "areyay"},
	{"quiet", "ietquay"},
	{"yellow", "ellowyay"},
	{"qu", "quay"},
	{"it's", "it'syay"},
	{".,()-#@!?[]", ".,()-#@!?[]"},
	{"test@test.com", "test@test.com"},
	{"LittleBig", "IttleBiglay"},
	{"I'm", "I'myay"},
	{"cyanide", "yanidecay"},
}

func TestNewEnglish(t *testing.T) {
	t.Run("constructor", func(t *testing.T) {
		tr := NewEnglish()
		givenType := reflect.TypeOf(*tr)
		wantedType := reflect.TypeOf((*English)(nil)).Elem()
		if givenType != wantedType {
			t.Errorf("got %q, want %q", givenType, wantedType)
		}
	})
}

func TestTranslate(t *testing.T) {
	translator := NewEnglish()
	for _, tt := range samples {
		t.Run(tt.in, func(t *testing.T) {
			s := translator.Translate(tt.in)
			if s != tt.out {
				t.Errorf("tried %q, got %q, want %q", tt.in, s, tt.out)
			}
		})
	}
}
