package db

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

var instance *sql.DB
var once sync.Once

func GetPostgresDb() *sql.DB {
	once.Do(func() {
		db, err := sql.Open(PostgresDriver, DataSourceName)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("Accessing %s ... ", DbName)

		fmt.Println("Connected!")

		// TOOD ver lugar melhor pra criar tabelas
		_, err = db.Exec(InitScript)
		if err != nil {
			fmt.Print(err.Error())
			panic(err.Error())
		}
		fmt.Println("Tables created")

		// guarda a instância para reuso
		instance = db
	})
	return instance
}

var InitScript = `
			CREATE TABLE IF NOT EXISTS balance (
				id VARCHAR(36) PRIMARY KEY,
				value VARCHAR(255) NOT NULL
			);

			CREATE TABLE IF NOT EXISTS usuario (
				id VARCHAR(36) PRIMARY KEY,
				cpf VARCHAR(14) NOT NULL UNIQUE,
				balance_id VARCHAR(36),
				CONSTRAINT fk_balance FOREIGN KEY (balance_id) REFERENCES balance(id),
				CONSTRAINT cpf_not_empty CHECK (cpf <> '')
			);
		`
