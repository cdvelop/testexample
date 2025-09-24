package testexample

import "testing"

func TestSaludar(t *testing.T) {

	result := Saludar("maria mercedes")

	expectativa := "Hello, maria mercedes!"

	if result != expectativa {
		t.Errorf("\nresultado: = %q;\nse esperaba: %q", result, expectativa)
	}

}
