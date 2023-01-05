package main

import (
	"log"

	"github.com/LouisMatos/challenge-backend-2-go/app/config"
	"github.com/LouisMatos/challenge-backend-2-go/app/database"
	"github.com/LouisMatos/challenge-backend-2-go/app/routes"
)

func main() {

	config.LoadAppConfig()

	log.Println("Iniciando conexão com o banco de dados")

	database.ConexaoComBancoDados(config.AppConfig.ConnectionString)

	database.Migrate()

	log.Println("Iniciando o servidor!")

	routes.HandleRequest(config.AppConfig.Port)

}
