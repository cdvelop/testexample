package sqlite

import "fmt"

type sqliteEngine struct {
	PathBb string
}

// config ruta archivo ej: .dbConfig
func New(config string) sqliteEngine {
	return sqliteEngine{
		PathBb: config,
	}
}

func (db sqliteEngine) Conectar() error {
	fmt.Println("conectando a base de datos Sqlite")

	return nil
}
