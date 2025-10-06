package controller

import (
	"AlderFurtado/BankGo.git/infra/cache"
	controller_error_handling "AlderFurtado/BankGo.git/internal/controller/error"
	"AlderFurtado/BankGo.git/internal/controller/util"
	"AlderFurtado/BankGo.git/internal/domain/usecase"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type GetUserBalanceHandler struct {
	GetUserBalanceValue usecase.GetUserBalanceValue
}

func NewGetUserBalanceValue(gubv usecase.GetUserBalanceValue) GetUserBalanceHandler {
	return GetUserBalanceHandler{GetUserBalanceValue: gubv}
}

func (gubc GetUserBalanceHandler) Invoke(c *gin.Context) {
	cpf := c.Query("cpf")
	//TODO error de cpf vazio
	redisKey := "cpf-balance-" + cpf
	balance, err := cache.GetRedisCache().Get(redisKey)
	log.Printf("Recebi request para CPF=%v e do receber do redis(%v) o valor %v", cpf, redisKey, balance)
	log.Printf("err =%v", err)
	if balance != "" {
		c.JSON(http.StatusOK, util.FormatJsonResponseHttp(balance, nil))
		return
	}

	value, err := gubc.GetUserBalanceValue.Invoke(cpf)
	log.Printf("Foi no banco")
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			util.FormatJsonResponseHttp(nil, controller_error_handling.ControllerInternal))
		return
	}
	log.Printf("Retorna do redis(%v) o valor = %v", redisKey, value)

	cache.GetRedisCache().Set(redisKey, value, time.Minute*5)
	c.JSON(http.StatusOK, util.FormatJsonResponseHttp(value, nil))
}

func (gubc GetUserBalanceHandler) InvokeWithoutCache(c *gin.Context) {
	cpf := c.Query("cpf")
	//TODO error de cpf vazio

	value, err := gubc.GetUserBalanceValue.Invoke(cpf)

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			util.FormatJsonResponseHttp(nil, controller_error_handling.ControllerInternal))
		return
	}
	c.JSON(http.StatusOK, util.FormatJsonResponseHttp(value, nil))
}
