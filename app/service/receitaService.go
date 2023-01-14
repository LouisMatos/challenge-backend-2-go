package service

import (
	"log"
	"strconv"
	"time"

	"github.com/LouisMatos/challenge-backend-2-go/app/model"
	"github.com/LouisMatos/challenge-backend-2-go/app/repository"
	"github.com/LouisMatos/challenge-backend-2-go/app/utils"
)

type ReceitaService interface {
	Save(receita model.ReceitaDTO) (model.Receita, bool)
	GetAll(descricao string) []model.Receita
	FindById(id string) model.Receita
	FindReceitaByAnoAndMes(ano string, mes string) []model.Receita
	Update(receita model.ReceitaDTO, id string) (model.Receita, bool)
	Delete(id string)
	ValidarReceitaJaExiste(Descricao string, Data time.Time) bool
}

type receitaService struct {
	receitaRepository repository.ReceitaRepository
}

func NewReceitaService(repo repository.ReceitaRepository) ReceitaService {
	return &receitaService{
		receitaRepository: repo,
	}
}

func (service *receitaService) Save(receitaDTO model.ReceitaDTO) (model.Receita, bool) {

	date, _ := time.Parse("02/01/2006 15:04:05", receitaDTO.Data+" 00:00:00")

	value, _ := strconv.ParseFloat(receitaDTO.Valor, 64)

	log.Println("Convertendo dto para objeto a ser salvo no banco de dados!")

	receita := model.Receita{
		Descricao: receitaDTO.Descricao,
		Data:      date,
		Valor:     utils.RoundUp(float64(value), 2),
	}

	isSaved := service.ValidarReceitaJaExiste(receita.Descricao, receita.Data)

	if !isSaved {

		receita = service.receitaRepository.Save(receita)

		log.Println("Receita salva no banco de dados!")

		return receita, false

	}

	return receita, true
}

func (service *receitaService) Update(receitaDTO model.ReceitaDTO, id string) (model.Receita, bool) {

	date, _ := time.Parse("02/01/2006 15:04:05", receitaDTO.Data+" 00:00:00")

	value, _ := strconv.ParseFloat(receitaDTO.Valor, 64)

	u, err := strconv.ParseUint(id, 0, 64)
	if err != nil {
		panic(err)
	}

	log.Println("Convertendo dto para objeto a ser atualizado no banco de dados!")

	receita := model.Receita{
		ID:        uint(u),
		Descricao: receitaDTO.Descricao,
		Data:      date,
		Valor:     utils.RoundUp(float64(value), 2),
	}

	isSaved := service.ValidarReceitaJaExiste(receita.Descricao, receita.Data)

	if !isSaved {

		receita = service.receitaRepository.Save(receita)

		log.Println("Receita atualizada no banco de dados!")

		return receita, false
	} else {
		return receita, true
	}
}

func (service *receitaService) FindReceitaByAnoAndMes(ano string, mes string) []model.Receita {

	if len(mes) == 1 {
		mes = "0" + mes
	}

	receitas := service.receitaRepository.FindReceitaByAnoAndMes(ano, mes)

	return receitas
}

func (service *receitaService) Delete(id string) {
	service.receitaRepository.Delete(id)
}

func (service *receitaService) ValidarReceitaJaExiste(descricao string, data time.Time) bool {

	receita := service.receitaRepository.AlreadyRegistered(descricao, data)

	if receita.ID == 0 {
		log.Println("Receita ainda não foi cadastrada!")
		return false
	} else {
		log.Println("Receita já foi cadastrada!")
		return true
	}

}

func (service *receitaService) FindById(id string) model.Receita {

	receita := service.receitaRepository.FindById(id)

	return receita
}

func (service *receitaService) GetAll(descricao string) []model.Receita {

	var receitas []model.Receita

	receitas = service.receitaRepository.GetAll(descricao, receitas)

	return receitas

}
