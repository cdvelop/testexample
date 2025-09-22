// test caja negra
package testexample_test

import (
	"testing"

	"github.com/cdvelop/testexample"
)

func TestUserInfoBlackBox(t *testing.T) {

	user := testexample.NewUser("Cesar", "cesar@example.com")

	result := user.GetInfo()

	expectativa := "Cesar <cesar@example.com>"

	if result != expectativa {
		t.Errorf("Error GenerarSaludo, se esperaba %q y se obtuvo %q", expectativa, result)
	}

}
