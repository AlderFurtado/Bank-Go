package main

import "AlderFurtado/BankGo.git/internal/route"

func main() {
	route.GetRoute().RunAuthApi(":4000")
}
