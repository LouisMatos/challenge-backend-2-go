package service

import (
	"log"
	"strconv"
	"time"

	"github.com/LouisMatos/challenge-backend-2-go/app/enum"
	"github.com/LouisMatos/challenge-backend-2-go/app/model"
	"github.com/LouisMatos/challenge-backend-2-go/app/repository"
	"github.com/LouisMatos/challenge-backend-2-go/app/utils"
)

type DespesaService interface {
	Save(despesaDTO model.DespesaDTO) (model.Despesa, bool)
	Update(despesaDTO model.DespesaDTO, id string) (model.Despesa, bool)
	FindDespesaByAnoAndMes(ano string, mes string) []model.Despesa
	GetAll(descricao string) []model.Despesa
	FindById(id string) model.Despesa
	Delete(id string)
	ValidarDespesaJaExiste(Descricao string, Data time.Time) bool
}

type despesaService struct {
	despesaRepository repository.DespesaRepository
}

func NewDespesaService(repo repository.DespesaRepository) DespesaService {
	return &despesaService{
		despesaRepository: repo,
	}
}

func (service *despesaService) Save(despesaDTO model.DespesaDTO) (model.Despesa, bool) {

	date, _ := time.Parse("02/01/2006 15:04:05", despesaDTO.Data+" 00:00:00")

	value, _ := strconv.ParseFloat(despesaDTO.Valor, 64)

	log.Println("Convertendo dto para objeto a ser salvo no banco de dados!")

	despesa := model.Despesa{
		Descricao: despesaDTO.Descricao,
		Data:      date,
		Valor:     utils.RoundUp(float64(value), 2),
		Categoria: verificaCategoria(despesaDTO.Categoria),
	}

	isSaved := service.ValidarDespesaJaExiste(despesa.Descricao, despesa.Data)

	if !isSaved {

		despesa = service.despesaRepository.Save(despesa)

		log.Println("Despesa salva no banco de dados!")

		return despesa, false

	} else {
		return despesa, true
	}

}

func (service *despesaService) ValidarDespesaJaExiste(descricao string, data time.Time) bool {

	despesa := service.despesaRepository.AlreadyRegistered(descricao, data)

	if despesa.ID == 0 {
		log.Println("Despesa ainda não foi cadastrada!")
		return false
	} else {
		log.Println("Despesa já foi cadastrada!")
		return true
	}

}

func (service *despesaService) GetAll(descricao string) []model.Despesa {

	var despesas []model.Despesa

	despesas = service.despesaRepository.GetAll(descricao, despesas)

	return despesas

}

func (service *despesaService) FindById(id string) model.Despesa {

	receita := service.despesaRepository.FindById(id)

	return receita

}

func (service *despesaService) Update(despesaDTO model.DespesaDTO, id string) (model.Despesa, bool) {

	date, _ := time.Parse("02/01/2006 15:04:05", despesaDTO.Data+" 00:00:00")

	value, _ := strconv.ParseFloat(despesaDTO.Valor, 64)

	u, err := strconv.ParseUint(id, 0, 64)
	if err != nil {
		panic(err)
	}

	log.Println("Convertendo dto para objeto a ser atualizado no banco de dados!")

	despesa := model.Despesa{
		ID:        uint(u),
		Descricao: despesaDTO.Descricao,
		Data:      date,
		Valor:     utils.RoundUp(float64(value), 2),
		Categoria: verificaCategoria(despesaDTO.Categoria),
	}

	isSaved := service.ValidarDespesaJaExiste(despesa.Descricao, despesa.Data)

	if !isSaved {

		despesa = service.despesaRepository.Update(despesa)

		log.Println("Despesa atualizada no banco de dados!")

		return despesa, false
	} else {
		return despesa, true
	}

}

func (service *despesaService) Delete(id string) {
	service.despesaRepository.Delete(id)
}

func (service *despesaService) FindDespesaByAnoAndMes(ano string, mes string) []model.Despesa {

	if len(mes) == 1 {
		mes = "0" + mes
	}

	receitas := service.despesaRepository.FindDespesaByAnoAndMes(ano, mes)

	return receitas
}

func verificaCategoria(categoria string) enum.Categoria {

	switch categoria {
	case "1":
		return 1
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	default:
		return 8
	}

}
