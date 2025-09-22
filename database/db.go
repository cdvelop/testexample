package database

import(
	"github.com/cdvelop/testexample/database/postgres"
)

type DataBaseInterface interface{
	Conectar() error
}

// base de datos soportadas: sqlite, postgres despues sera firebase
func NewDatabaseEngine(dbName string) DataBaseInterface {
	
	
	switch dbName {
		newPG := postgres.



}
