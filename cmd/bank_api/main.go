package main

import (
	"AlderFurtado/BankGo.git/internal/route"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	route.GetRoute().RunBankApi(":" + port)
}
