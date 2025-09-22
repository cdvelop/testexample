// test caja negra
package testexample_test

import (
	"testing"

	"github.com/cdvelop/testexample/database"
	//"github.com/cdvelop/testexample/userInterface"
)

type dbInterfaceTest interface {
	Conectar() error
}

func TestIntegration(t *testing.T) {

	testData := []struct {
		testName      string
		dbInputType   string
		expectedError string
	}{
		{testName: "Base de datos se espera error", dbInputType: "post", expectedError: "db not found"},
		{testName: "Base de datos SQLite ok", dbInputType: "sqlite", expectedError: "connection ok"},
		{testName: "Base de datos Postgres ok", dbInputType: "postgres", expectedError: "connection ok"},
	}

	for _, tt := range testData {
		t.Run(tt.dbInputType, func(t *testing.T) {
			_, err := database.NewDatabaseEngine(tt.dbInputType)
			errorEsperado := "connection ok"

			if err != nil {
				// transformo el error a string
				errorEsperado = err.Error()
			}

			if errorEsperado != tt.expectedError {
				t.Fatalf("Error esperado: %v, obtenido: %v", tt.expectedError, errorEsperado)
			}

		})
	}

}
