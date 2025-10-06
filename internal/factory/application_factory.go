package factory

import (
	"AlderFurtado/BankGo.git/infra/db"
	transactionController "AlderFurtado/BankGo.git/internal/controller/transaction"
	useController "AlderFurtado/BankGo.git/internal/controller/user"
	"AlderFurtado/BankGo.git/internal/data/repository"
	"AlderFurtado/BankGo.git/internal/domain/usecase"
)

func getUserRepositoryImpl() repository.UserRepositoryImpl {
	return repository.NewUserRepositoryImpl(db.GetPostgresDb())
}

func GetCreateUserUseCase() usecase.CreateUser {
	return *usecase.NewCreateUser(getUserRepositoryImpl())
}

func GetGetUserBalanceValueUseCase() usecase.GetUserBalanceValue {
	return *usecase.NewGetUserBalanceValue(getUserRepositoryImpl())
}

func GetCreateUserHandler() useController.CreateUserHandler {
	return useController.NewCreateUserController(GetCreateUserUseCase())
}

func GetGetUserBalanceHandler() useController.GetUserBalanceHandler {
	return useController.NewGetUserBalanceValue(GetGetUserBalanceValueUseCase())
}

func GetTransferHandler() transactionController.TransferHandler {
	return transactionController.NewTransferHandler()
}
