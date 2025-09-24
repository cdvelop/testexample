package testexample

import (
	"testing"
)

/*func TestExampleFibonacci(t *testing.T) {

	result := Fibonacci(5)

	expectativa := 8

	if result != expectativa {
		fmt.Println("Error Fibonacci, se esperaba " + fmt.Sprint(expectativa) + " y se obtuvo " + fmt.Sprint(result))
	}

}*/

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name        string // nombre del caso
		num         int    // datos de entrada
		expectativa int    // expectativa
	}{
		{name: "Caso 1", num: 5, expectativa: 8},
		{name: "Caso 2", num: 0, expectativa: 1},
		{name: "Caso 3", num: 1, expectativa: 1},
		{name: "Caso 4", num: 2, expectativa: 2},
		{name: "Caso 5", num: 3, expectativa: 3},
		{name: "Caso 6", num: 4, expectativa: 5},
		{name: "Caso 7", num: 6, expectativa: 13},
		{name: "Caso 8", num: 7, expectativa: 21},
		{name: "Caso 9", num: 8, expectativa: 34},
		{name: "Caso 10", num: 9, expectativa: 55},

		// TODO: agregar al menos 10 casos
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Fibonacci(tt.num)
			if got != tt.expectativa {
				t.Errorf("Fibonacci(%d) = %d; want %d", tt.num, got, tt.expectativa)
			}
		})
	}
}
