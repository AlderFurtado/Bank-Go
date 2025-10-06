package repository

import (
	postgresdb "AlderFurtado/BankGo.git/infra/db"
	repository_error_handling "AlderFurtado/BankGo.git/internal/data/error"
	"AlderFurtado/BankGo.git/internal/domain/entity"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

//TODO melhorar teste

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	assert.NoError(t, err)

	// cria tabelas
	_, err = db.Exec(postgresdb.InitScript)
	assert.NoError(t, err)

	return db
}

func TestUserRepository_Save_UserValid(t *testing.T) {

	db := setupTestDB(t)
	repo := NewUserRepositoryImpl(db)

	user := entity.User{Cpf: "12345678900", Balance: entity.Balance{Value: "100.0"}}
	err := repo.Save(user)
	assert.NoError(t, err)

	var cpf string
	row := db.QueryRow("SELECT cpf FROM usuario LIMIT 1")
	err = row.Scan(&cpf)
	assert.NoError(t, err)
	assert.Equal(t, "12345678900", cpf)
}

func TestUserRepository_Save_UserWithCpfInvalid(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepositoryImpl(db)

	user := entity.User{Cpf: "", Balance: entity.Balance{Value: "100.0"}}
	repo.Save(user)

	count := 0
	rows, _ := db.Query("SELECT cpf FROM usuario LIMIT 1")
	for rows.Next() {
		count++
	}
	assert.Equal(t, 0, count)
}

func TestUserRepository_Save_DbInternalError(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepositoryImpl(db)

	user := entity.User{Cpf: "", Balance: entity.Balance{Value: "100.0"}}
	err := repo.Save(user)
	assert.Equal(t, repository_error_handling.DbInternalError.Error(), err.Error())
}
