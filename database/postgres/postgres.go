package postgres

import "fmt"

type postgresEngine struct {
	pathBb string
}

// config cadena de conexion ej: posgres:xxx..
func New(config string) postgresEngine {

	return postgresEngine{
		pathBb: config,
	}
}

func (db postgresEngine) Conectar() error {
	fmt.Println("conectando a base de datos Postgres")

	return nil
}
