package model

type Resumo struct {
	ValorTotalReceita   float64 `json:"valorTotalReceita"`
	ValorTotalDespesa   float64 `json:"valorTotalDespesa"`
	SaldoFinal          float64 `json:"saldoFinal"`
	ValorTotalCategoria string  `json:"valorTotalCategoria"`
}

type CategoriaResumo struct {
	Valor     float64 `json:"valor"`
	Categoria string  `json:"categoria"`
}
