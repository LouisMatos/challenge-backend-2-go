package repository

import (
	"time"

	"github.com/LouisMatos/challenge-backend-2-go/app/model"
	"gorm.io/gorm"
)

type DespesaRepository interface {
	Save(despesa model.Despesa) model.Despesa
	Update(despesa model.Despesa) model.Despesa
	GetAll(descricao string, despesas []model.Despesa) []model.Despesa
	FindDespesaByAnoAndMes(ano string, mes string) []model.Despesa
	FindById(id string) model.Despesa
	Delete(id string)
	AlreadyRegistered(descricao string, data time.Time) model.Despesa
}

type despesaRepository struct {
	db *gorm.DB
}

func NewDespesaRepository(dbConnection *gorm.DB) DespesaRepository {
	return &despesaRepository{
		db: dbConnection,
	}
}

func (repository *despesaRepository) Save(despesa model.Despesa) model.Despesa {

	repository.db.Create(&despesa)

	return despesa
}

func (repository *despesaRepository) Update(despesa model.Despesa) model.Despesa {

	repository.db.Save(&despesa)

	return despesa
}

func (repository *despesaRepository) AlreadyRegistered(descricao string, data time.Time) model.Despesa {

	var despesa model.Despesa

	repository.db.Where("descricao ILIKE ? AND TO_CHAR(data, 'yyyy-mm') LIKE ?", descricao, data.Format("2006-01")).Find(&despesa)

	return despesa
}

func (repository *despesaRepository) GetAll(descricao string, despesas []model.Despesa) []model.Despesa {

	repository.db.Where("descricao ILIKE ?", "%"+descricao+"%").Find(&despesas)

	return despesas
}

func (repository *despesaRepository) FindById(id string) model.Despesa {

	var despesa model.Despesa

	repository.db.First(&despesa, id)

	return despesa
}

func (repository *despesaRepository) Delete(id string) {

	var despesa model.Despesa

	repository.db.Delete(&despesa, id)
}

func (repository *despesaRepository) FindDespesaByAnoAndMes(ano string, mes string) []model.Despesa {

	var despesas []model.Despesa

	repository.db.Where("(TO_CHAR(data, 'YYYY-MM')) = ?", ""+ano+"-"+mes).Find(&despesas)

	return despesas

}
