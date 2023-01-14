package service

import (
	"github.com/LouisMatos/challenge-backend-2-go/app/enum"
	"github.com/LouisMatos/challenge-backend-2-go/app/model"
	"github.com/LouisMatos/challenge-backend-2-go/app/repository"
	"github.com/LouisMatos/challenge-backend-2-go/app/utils"
)

type ResumoService interface {
	GetMonthSummary(ano string, mes string) model.Resumo
}

type resumoService struct {
	resumoRepository repository.ResumoRepository
}

func NewResumoService(repo repository.ResumoRepository) ResumoService {
	return &resumoService{
		resumoRepository: repo,
	}
}

func (service *resumoService) GetMonthSummary(ano string, mes string) model.Resumo {

	var resumo model.Resumo

	var despesas []model.Despesa
	var receitas []model.Receita

	valorCategoria := make([]model.CategoriaResumo, 8)

	if len(mes) == 1 {
		mes = "0" + mes
	}

	despesas = service.resumoRepository.GetAllDespesasByAnoAndMes(ano, mes)

	receitas = service.resumoRepository.GetAllReceitasByAnoAndMes(ano, mes)

	for i := 0; i < len(despesas); i++ {
		resumo.ValorTotalDespesa = resumo.ValorTotalDespesa + float64(despesas[i].Valor)
		realizaSuamarizacaoGastoCategoria(valorCategoria, despesas[i].Valor, despesas[i].Categoria)
	}

	for i := 0; i < len(receitas); i++ {
		resumo.ValorTotalReceita = resumo.ValorTotalDespesa + float64(receitas[i].Valor)
	}

	resumo.ValorTotalDespesa = utils.RoundUp(resumo.ValorTotalDespesa, 2)
	resumo.ValorTotalReceita = utils.RoundUp(resumo.ValorTotalReceita, 2)
	resumo.SaldoFinal = utils.RoundUp(resumo.ValorTotalReceita-resumo.ValorTotalDespesa, 2)
	resumo.ValorTotalCategoria = valorCategoria

	for i := 0; i < len(valorCategoria); i++ {
		valorCategoria[i].Valor = utils.RoundUp(valorCategoria[i].Valor, 2)
	}

	return resumo

}

func realizaSuamarizacaoGastoCategoria(resumo []model.CategoriaResumo, valor float64, categoria enum.Categoria) {

	switch categoria.EnumIndex() {

	case 1:
		resumo[categoria.EnumIndex()-1].Valor += valor
		resumo[categoria.EnumIndex()-1].Categoria = categoria.String()
	case 2:
		resumo[categoria.EnumIndex()-1].Valor += valor
		resumo[categoria.EnumIndex()-1].Categoria = categoria.String()
	case 3:
		resumo[categoria.EnumIndex()-1].Valor += valor
		resumo[categoria.EnumIndex()-1].Categoria = categoria.String()
	case 4:
		resumo[categoria.EnumIndex()-1].Valor += valor
		resumo[categoria.EnumIndex()-1].Categoria = categoria.String()
	case 5:
		resumo[categoria.EnumIndex()-1].Valor += valor
		resumo[categoria.EnumIndex()-1].Categoria = categoria.String()
	case 6:
		resumo[categoria.EnumIndex()-1].Valor += valor
		resumo[categoria.EnumIndex()-1].Categoria = categoria.String()
	case 7:
		resumo[categoria.EnumIndex()-1].Valor += valor
		resumo[categoria.EnumIndex()-1].Categoria = categoria.String()
	case 8:
		resumo[categoria.EnumIndex()-1].Valor += valor
		resumo[categoria.EnumIndex()-1].Categoria = categoria.String()
	}
}
