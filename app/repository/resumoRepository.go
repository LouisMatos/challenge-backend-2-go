package repository

import (
	"github.com/LouisMatos/challenge-backend-2-go/app/model"
	"gorm.io/gorm"
)

type ResumoRepository interface {
	GetAllDespesasByAnoAndMes(ano string, mes string) []model.Despesa
	GetAllReceitasByAnoAndMes(ano string, mes string) []model.Receita
}

type resumoRepository struct {
	db *gorm.DB
}

func NewResumoRepository(dbConnection *gorm.DB) ResumoRepository {
	return &resumoRepository{
		db: dbConnection,
	}
}

func (repository *resumoRepository) GetAllDespesasByAnoAndMes(ano string, mes string) []model.Despesa {

	var despesas []model.Despesa

	repository.db.Where("(TO_CHAR(data, 'YYYY-MM')) = ?", ""+ano+"-"+mes).Find(&despesas)

	return despesas
}

func (repository *resumoRepository) GetAllReceitasByAnoAndMes(ano string, mes string) []model.Receita {

	var receitas []model.Receita

	repository.db.Where("(TO_CHAR(data, 'YYYY-MM')) = ?", ""+ano+"-"+mes).Find(&receitas)

	return receitas
}
