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

func (db sqliteEngine) Conectar() {
	fmt.Println("conectando a base de datos Sqlite")
}
