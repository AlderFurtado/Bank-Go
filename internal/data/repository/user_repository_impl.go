package repository

import (
	repository_error_handling "AlderFurtado/BankGo.git/internal/data/error"
	"AlderFurtado/BankGo.git/internal/domain/entity"
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// TODO Impl para Sql
type UserRepositoryImpl struct {
	Db *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) UserRepositoryImpl {
	return UserRepositoryImpl{Db: db}
}

func (uri UserRepositoryImpl) Save(user entity.User) error {
	db := uri.Db
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return repository_error_handling.DbInternalError
	}

	idBalance := uuid.New()
	rand.Seed(time.Now().UnixNano())

	_, err = tx.Exec("INSERT INTO balance (id, value) VALUES ($1, $2);", idBalance, user.Balance.Value)
	if err != nil {
		tx.Rollback() // desfaz se deu erro
		fmt.Println("Erro ao inserir balance:", err)
		return repository_error_handling.DbInternalError
	}

	idUser := uuid.New()
	fmt.Println(user.Cpf)
	_, err = tx.Exec("INSERT INTO usuario (id, cpf, balance_id) VALUES ($1, $2, $3);", idUser, user.Cpf, idBalance)
	if err != nil {
		tx.Rollback() // desfaz se deu erro
		fmt.Println("Erro ao inserir usuario:", err)
		return repository_error_handling.DbInternalError
	}
	if err := tx.Commit(); err != nil {
		fmt.Println("Erro ao dar commit:", err)
		return repository_error_handling.DbInternalError
	}
	return nil
}

func (uri UserRepositoryImpl) GetByCpf(cpf string) (entity.User, error) {
	db := uri.Db
	rows, err := db.Query(`SELECT b.value
			FROM usuario u
			JOIN balance b ON u.balance_id = b.id
			WHERE u.cpf = $1;`, cpf)
	if err != nil {
		fmt.Println("Erro ao encontrar registro:", err)
		return entity.User{}, repository_error_handling.DbInternalError
	}
	balance := ""
	for rows.Next() {
		if err := rows.Scan(&balance); err != nil {
			fmt.Println("Erro ao inserir balance:", err)
			return entity.User{}, repository_error_handling.DbInternalError
		}
	}
	if balance == "" {
		fmt.Println("Erro ao encontrar registro:", err)
		return entity.User{}, repository_error_handling.DbNotFound
	}
	return entity.User{Cpf: "00000000", Balance: entity.Balance{Value: balance}}, nil
}
