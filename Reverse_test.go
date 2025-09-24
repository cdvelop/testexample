package testexample

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Palabra simple", "hola", "aloh"},
		{"Palabra con tilde", "ábaco", "ocabá"},
		{"Cadena vacía", "", ""},
		{"Un solo carácter", "a", "a"},
		{"Dos caracteres", "ab", "ba"},
		{"Unicode", "¡Go!", "!oG¡"},
		{"Emoji", "😀👍", "👍😀"},
		{"Frase", "anita lava la tina", "anit al aval atina"},
		{"Espacios", " a b ", " b a "},
		{"Números", "12345", "54321"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reverse(tt.input)
			if got != tt.expected {
				t.Errorf("Reverse(%q) = %q; want %q", tt.input, got, tt.expected)
			}
		})
	}
}
