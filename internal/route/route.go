package route

import (
	"net/http"
	"os"
	"sync"

	"AlderFurtado/BankGo.git/internal/factory"

	"github.com/gin-gonic/gin"
)

type route struct {
	Engine *gin.Engine
}

var instance *route
var once sync.Once

func GetRoute() *route {
	engine := gin.Default()
	once.Do(func() {
		instance = &route{Engine: engine}
	})
	return instance
}

func addHeaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Antes de chamar o handler
		c.Writer.Header().Set("X-App-Name", os.Getenv("API_NAME")) // header customizado

		c.Next() // chama o próximo handler

		// Depois do handler, você pode adicionar mais headers se quiser
		c.Writer.Header().Set("X-Powered-By", "Gin")
	}
}

func (r *route) RunBankApi(port string) {

	GetRoute().Engine.Use(addHeaderMiddleware())
	GetRoute().Engine.POST("/user/create_user", factory.GetCreateUserHandler().Invoke)
	GetRoute().Engine.GET("/user/balance", factory.GetGetUserBalanceHandler().Invoke)
	GetRoute().Engine.GET("/user/balance_no_cache", factory.GetGetUserBalanceHandler().InvokeWithoutCache)
	GetRoute().Engine.POST("/transfer", factory.GetTransferHandler().Invoke)

	GetRoute().Engine.Run(port)
}

func (r *route) RunAuthApi(port string) {
	GetRoute().Engine.GET("/validate", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"ok": "ok"})
	})
	GetRoute().Engine.GET("/generate", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"X-API-Key": "MINHA_CHAVE_SECRETA"})
	})
	GetRoute().Engine.Run(port)
}
