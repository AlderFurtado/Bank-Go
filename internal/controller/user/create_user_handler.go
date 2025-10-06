package controller

import (
	controller_error_handling "AlderFurtado/BankGo.git/internal/controller/error"
	"AlderFurtado/BankGo.git/internal/controller/util"
	"AlderFurtado/BankGo.git/internal/domain/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserHandler struct {
	CreateUserUseCase usecase.CreateUser
}

// TODO Teste
type createUserHandlerDto struct {
	Cpf string "json:cpf"
}

func NewCreateUserController(cuuc usecase.CreateUser) CreateUserHandler {
	return CreateUserHandler{CreateUserUseCase: cuuc}
}

func (cuc CreateUserHandler) Invoke(c *gin.Context) {
	var req createUserHandlerDto
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.FormatJsonResponseHttp(nil, controller_error_handling.ControllerInputInvalid))
		return
	}

	err := cuc.CreateUserUseCase.Invoke(req.Cpf)

	if err != nil {
		c.JSON(http.StatusInternalServerError, util.FormatJsonResponseHttp(nil, controller_error_handling.ControllerInternal))
		return
	}
	c.JSON(http.StatusOK, util.FormatJsonResponseHttp("ok", nil))
}
