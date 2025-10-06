package transaction

import (
	"encoding/json"
	"fmt"
	"net/http"

	messagebroker "AlderFurtado/BankGo.git/infra/message_broker"
	controller_error_handling "AlderFurtado/BankGo.git/internal/controller/error"
	"AlderFurtado/BankGo.git/internal/controller/util"

	"github.com/gin-gonic/gin"
)

type TransferHandler struct{}

type transferHandlerDto struct {
	Cpf   string `json:"cpf"`
	Value int    `json:"value"`
}

func NewTransferHandler() TransferHandler {
	return TransferHandler{}
}

func (th TransferHandler) Invoke(c *gin.Context) {
	var req transferHandlerDto
	if err := c.BindJSON(&req); err != nil {
		fmt.Print(err.Error())
		c.JSON(http.StatusBadRequest, util.FormatJsonResponseHttp(
			nil,
			controller_error_handling.ControllerInputInvalid,
		))
		return
	}

	// Converte struct para JSON
	jsonBytes, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Transforma []byte em string
	jsonStr := string(jsonBytes)
	err = messagebroker.Producer("test", jsonStr)
	if err != nil {
		fmt.Print(err.Error())
		c.JSON(http.StatusBadRequest, util.FormatJsonResponseHttp(
			nil,
			controller_error_handling.ControllerInternal,
		))
		return
	}

	c.JSON(http.StatusOK, util.FormatJsonResponseHttp("Transfer is processing", nil))
}
