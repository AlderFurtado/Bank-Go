package repository

import "AlderFurtado/BankGo.git/internal/domain/entity"

type UserRepository interface {
	Save(user entity.User) error
	GetByCpf(cpf string) (entity.User, error)
	// Update(ursi entity.User) entity.User
	// Delete(id string) error
}
