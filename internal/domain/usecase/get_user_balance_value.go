package usecase

import "AlderFurtado/BankGo.git/internal/domain/repository"

type GetUserBalanceValue struct {
	repository repository.UserRepository
}

func NewGetUserBalanceValue(repoistory repository.UserRepository) *GetUserBalanceValue {
	return &GetUserBalanceValue{repository: repoistory}
}

func (gubv GetUserBalanceValue) Invoke(cpf string) (string, error) {
	value, err := gubv.repository.GetByCpf(cpf)
	if err != nil {
		return "", err
	}
	return value.Balance.Value, nil
}
