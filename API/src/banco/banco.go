package banco

import (
	"api/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.StringConexaoBancoDeDados)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
