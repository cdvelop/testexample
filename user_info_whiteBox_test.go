package testexample

// test caja blanca
import (
	"testing"

	"github.com/cdvelop/testexample/database"
)

func TestDatabaseWhiteBox(t *testing.T) {

	dbConfig := database.NewDBConfig()

	dbConfig.IsLoginOKParaTest(true)

	if dbConfig.IsLoginOK() != true {
		t.Errorf("Error en NewDBConfig, se esperaba loginOK=true y se obtuvo loginOK=%v", dbConfig.IsLoginOK())
	}
}

func TestUserInfoWhiteBox(t *testing.T) {

	user := NewUser("Cesar", "cesar@example.com")

	result := user.GetInfo()

	expectativa := "Cesar <cesar@example.com>"

	if result != expectativa {
		t.Errorf("Error GenerarSaludo, se esperaba %q y se obtuvo %q", expectativa, result)
	}

}
