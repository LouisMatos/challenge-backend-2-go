package model

type Receita struct {
	ID        uint    `gorm:"primary_key" json:"id,omitempty"`
	Descricao string  `gorm:"not null" json:"descricao"`
	Valor     float64 `gorm:"not null" json:"valor,omitempty"`
	Data      string  `gorm:"not null" json:"data,omitempty"`
}
