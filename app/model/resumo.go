package model

type Resumo struct {
	ValorTotalReceita   float64           `json:"valor_total_receita"`
	ValorTotalDespesa   float64           `json:"valor_total_despesa"`
	SaldoFinal          float64           `json:"saldo_final"`
	ValorTotalCategoria []CategoriaResumo `json:"valor_total_categoria"`
}

type CategoriaResumo struct {
	Valor     float64 `json:"valor"`
	Categoria string  `json:"categoria"`
}
