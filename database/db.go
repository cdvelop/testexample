package database

type DBConfig struct {
	loginOK bool // campo no exportado
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		loginOK: false,
	}
}

func (db *DBConfig) IsLoginOKParaTest(set bool) bool {
	db.loginOK = set
	return db.loginOK
}

func (db *DBConfig) IsLoginOK() bool {
	// Aquí iría la lógica real de verificación de login
	return db.loginOK
}
