package database

import (
	"log"

	"github.com/LouisMatos/challenge-backend-2-go/app/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConexaoComBancoDados(connectionString string) {

	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados!")
	}
	log.Println("Conexão ao banco de dados realizado com sucesso!")

}

func Migrate() {
	DB.AutoMigrate(&model.Despesa{})
	DB.AutoMigrate(&model.Receita{})
	log.Println("Database Migration Completed...")
}
