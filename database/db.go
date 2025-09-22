package database

import (
	"fmt"

	"github.com/cdvelop/testexample/database/postgres"
	"github.com/cdvelop/testexample/database/sqlite"
)

type DataBaseInterface interface {
	Conectar()
}

// base de datos soportadas: "sqlite", "postgres" despues sera firebase
func NewDatabaseEngine(dbName string) DataBaseInterface {

	switch dbName {

	case "postgres":

		newPG := postgres.New("postgres://user:pass@localhost:5432/dbname")

		return newPG

	case "sqlite":

		newSqLite := sqlite.New("./midb")

		return newSqLite

	default:
		fmt.Println("NO existe esta db:", dbName)

		return nil
	}

}
