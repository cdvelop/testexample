package database

import (
	"fmt"

	"github.com/cdvelop/testexample/database/postgres"
	"github.com/cdvelop/testexample/database/sqlite"
)

type DataBaseInterface interface {
	Conectar() error
}

// base de datos soportadas: "sqlite", "postgres" despues sera firebase
func NewDatabaseEngine(dbName string) (DataBaseInterface, error) {

	switch dbName {

	case "postgres":

		newPG := postgres.New("postgres://user:pass@localhost:5432/dbname")

		return newPG, nil

	case "sqlite":

		newSqLite := sqlite.New("./midb")

		return newSqLite, nil

	default:

		return nil, fmt.Errorf("NO existe esta db:", dbName)
	}

}
