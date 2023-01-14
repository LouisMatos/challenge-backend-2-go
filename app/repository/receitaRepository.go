package repository

import (
	"time"

	"github.com/LouisMatos/challenge-backend-2-go/app/model"
	"gorm.io/gorm"
)

type ReceitaRepository interface {
	Save(receita model.Receita) model.Receita
	GetAll(descricao string, receitas []model.Receita) []model.Receita
	FindById(id string) model.Receita
	FindReceitaByAnoAndMes(ano string, mes string) []model.Receita
	Delete(id string)
	AlreadyRegistered(descricao string, data time.Time) model.Receita
}

type receitaRepository struct {
	db *gorm.DB
}

func NewReceitaRepository(dbConnection *gorm.DB) ReceitaRepository {
	return &receitaRepository{
		db: dbConnection,
	}
}

func (repository *receitaRepository) Save(receita model.Receita) model.Receita {

	repository.db.Create(&receita)

	return receita
}

func (repository *receitaRepository) GetAll(descricao string, receitas []model.Receita) []model.Receita {

	repository.db.Where("descricao ILIKE ?", "%"+descricao+"%").Find(&receitas)

	return receitas
}

func (repository *receitaRepository) AlreadyRegistered(descricao string, data time.Time) model.Receita {

	var receita model.Receita

	repository.db.Where("descricao ILIKE ? AND TO_CHAR(data, 'yyyy-mm') LIKE ?", descricao, data.Format("2006-01")).Find(&receita)

	return receita
}

func (repository *receitaRepository) Delete(id string) {

	var receita model.Receita

	repository.db.Delete(&receita, id)
}

func (repository *receitaRepository) FindById(id string) model.Receita {

	var receita model.Receita

	repository.db.First(&receita, id)

	return receita
}

func (repository *receitaRepository) FindReceitaByAnoAndMes(ano string, mes string) []model.Receita {

	var receitas []model.Receita

	repository.db.Where("(TO_CHAR(data, 'YYYY-MM')) = ?", ""+ano+"-"+mes).Find(&receitas)

	return receitas

}
