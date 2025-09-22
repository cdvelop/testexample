// test caja negra
package testexample_test

import (
	"testing"

	"github.com/cdvelop/testexample/database"
	"github.com/cdvelop/testexample/userInterface"
)

func TestIntegration(t *testing.T) {

	db, err := database.NewDatabaseEngine("posg")
	if err != nil {
		t.Failed()
	}

	err = db.Conectar()
	if err != nil {
		t.Fatalf("fallo el test de la db")
	}

	ui := userInterface.New("<form>inputxx</form>")

	formHtml := ui.Render()
	if formHtml == "" {
		t.Fatalf("fallo el test de la UI")
	}

}
