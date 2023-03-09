package banco

import (
	"api/src/commons/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver
)

func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.Uri)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}