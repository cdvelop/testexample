package testexample

import (
	"fmt"
	"testing"
)

func TestExampleFunGenerarSaludo(t *testing.T) {

	result := GenerarSaludo("Cesar")

	expectativa := "Hello, Cesar!"

	if result != expectativa {
		fmt.Println("Error GenerarSaludo, se esperaba " + expectativa + " y se obtuvo " + result)
	}

}
