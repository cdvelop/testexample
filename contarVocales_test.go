package testexample

import (
	"testing"
)

/*func TestExampleContarVocales(t *testing.T) {

	result := ContarVocales("Galletas")

	expectativa := 3

	if result != expectativa {
		fmt.Println("Error ContarVocales, se esperaba " + fmt.Sprint(expectativa) + " y se obtuvo " + fmt.Sprint(result))
	}
}*/

func TestContarVocales(t *testing.T) {
	tests := []struct {
		name     string
		word     string
		expected int
	}{
		{name: "caso1", word: "Gonzalo", expected: 3},
		{name: "caso2", word: "Agua", expected: 3},
		{name: "caso3", word: "xyz", expected: 0},
		{name: "caso4", word: "", expected: 0},
		{name: "caso5", word: "AEIOUaeiou", expected: 10},
		{name: "caso6", word: "bcdfghjklmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ", expected: 0},
		{name: "caso7", word: "Hello, World!", expected: 3},
		{name: "caso8", word: "¡Hola, mundo!", expected: 4},
		{name: "caso9", word: "12345", expected: 0},
		{name: "caso10", word: "paralelepipedo", expected: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContarVocales(tt.word)
			if got != tt.expected {
				t.Errorf("ContarVocales(%v) = %v; expected %v", tt.word, got, tt.expected)
			}
		})
	}
}
