package database

import (
	"log"
	"time"

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
	} else {
		dbConfig, _ := DB.DB()
		dbConfig.SetMaxOpenConns(25)
		dbConfig.SetMaxIdleConns(25)
		dbConfig.SetConnMaxLifetime(5 * time.Minute)
	}
	log.Println("Conexão ao banco de dados realizado com sucesso!")

}

func GetConnection() *gorm.DB {
	return DB
}

func Migrate() {
	DB.AutoMigrate(&model.Despesa{})
	DB.AutoMigrate(&model.Receita{})
	DB.AutoMigrate(&model.Usuario{})
	log.Println("Migração do banco de dados concluída...")
}
