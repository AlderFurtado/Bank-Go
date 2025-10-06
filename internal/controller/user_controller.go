// package controller

// import (
// 	controller_error_handling "AlderFurtado/BankGo.git/internal/controller/error"
// 	"AlderFurtado/BankGo.git/internal/factory"
// 	"AlderFurtado/BankGo.git/internal/route"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type UserRequest struct {
// 	CPF string `json:"cpf"`
// }

// // TODO melhorar essa parte
// func InitControllers() {
// 	route.GetRoute().Engine.POST("/user/create_user", func(c *gin.Context) {
// 		var req UserRequest
// 		if err := c.BindJSON(&req); err != nil {
// 			c.JSON(http.StatusBadRequest, formatJsonResponseHttp(nil, controller_error_handling.ControllerBodyInvalid))
// 			return
// 		}

// 		cuuc := factory.GetCreateUserUseCase()
// 		err := cuuc.Invoke(req.CPF)

// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, formatJsonResponseHttp(nil, controller_error_handling.ControllerInternal))
// 		}
// 		c.JSON(http.StatusOK, formatJsonResponseHttp("ok", nil))
// 	})

// 	route.GetRoute().Engine.GET("/user/balance", func(c *gin.Context) {
// 		cpf := c.Query("cpf")

// 		cuuc := factory.GetGetUserBalanceValueUseCase()
// 		value, err := cuuc.Invoke(cpf)

// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError,
// 				formatJsonResponseHttp(nil, controller_error_handling.ControllerInternal))
// 		}
// 		c.JSON(http.StatusOK, formatJsonResponseHttp(value, nil))
// 	})

// 	route.GetRoute().Engine.GET("health", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, formatJsonResponseHttp("ok", nil))
// 	})

// 	go route.GetRoute().Run(":3000")
// }

//	func formatJsonResponseHttp(result any, err error) map[string]any {
//		return gin.H{
//			"result": result,
//			"error":  err,
//		}
//	}
package controller
