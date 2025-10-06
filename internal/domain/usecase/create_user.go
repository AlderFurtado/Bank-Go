package usecase

import (
	"AlderFurtado/BankGo.git/internal/domain/entity"
	"AlderFurtado/BankGo.git/internal/domain/repository"
	"AlderFurtado/BankGo.git/internal/validation"
	"errors"
)

type CreateUser struct {
	repository repository.UserRepository
}

func NewCreateUser(repoistory repository.UserRepository) *CreateUser {
	return &CreateUser{repository: repoistory}
}

func (cu CreateUser) Invoke(cpf string) error {
	if !validation.ValidCpf(cpf) {
		return errors.New("Cpf invalid")
	}
	balance := entity.Balance{Value: "0"}
	user := entity.User{Cpf: cpf, Balance: balance}
	err := cu.repository.Save(user)
	if err != nil {
		return err
	}
	return nil
}
