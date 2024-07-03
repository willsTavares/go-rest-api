package main

import (
	"fmt"

	"github.com/willsTavares/go-rest-api/database"
	"github.com/willsTavares/go-rest-api/routes"
)

func main() {
	database.ConnectWithDatabase()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
