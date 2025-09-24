package testexample

import (
	"errors"
)

type Database interface {
	Conectar() error
}

type SQLite struct{}

func (s SQLite) Conectar() error { return nil }

type Postgres struct{}

func (p Postgres) Conectar() error { return nil }

func NewDatabaseEngine(dbType string) (Database, error) {
	switch dbType {
	case "sqlite":
		return SQLite{}, nil
	case "postgres":
		return Postgres{}, nil
	default:
		return nil, errors.New("db not found")
	}
}
