package translation

import "testing"

type testInterface struct{}

var translateMock func(text string) string

func (ti *testInterface) Translate(text string) string {
	return translateMock(text)
}

func TestRegister(t *testing.T) {
	t.Run("translator registration", func(t *testing.T) {
		translators = make(map[string]Translator)
		trMock := &testInterface{}
		Register("dummy", trMock)

		if translators == nil {
			t.Errorf("got %v, want []Translator", nil)
		}

		trTest, ok := translators["dummy"]
		if !ok {
			t.Error("translator was not registered")
		}
		if trTest != trMock {
			t.Errorf("got %t, want %t", trTest, trMock)
		}
	})

	t.Run("translator registration panic on nil", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("the code did not panic on nil registration")
			}
		}()

		Register("test", nil)
	})

	t.Run("translator registration panic on duplicate", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("the code did not panic on nil registration")
			}
		}()

		trMock := &testInterface{}
		Register("test", trMock)
		Register("test", trMock)
	})
}

func TestTranslator(t *testing.T) {
	t.Run("new parser", func(t *testing.T) {
		translators = make(map[string]Translator)
		trMock := &testInterface{}
		translators["dummy"] = trMock

		tr, err := NewTranslator("dummy")
		if err != nil {
			t.Fatal("can not instantiate translator")
		}

		if tr != trMock {
			t.Errorf("got %t, want %t", tr, trMock)
		}
	})

	t.Run("no translator for the lang key", func(t *testing.T) {
		translators = make(map[string]Translator)
		tr, err := NewTranslator("dummy")
		if err == nil {
			t.Errorf("got nil, want error")
		}
		if tr != nil {
			t.Errorf("got %t, want nil", tr)
		}
	})
}
