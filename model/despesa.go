package model

import "gorm.io/gorm"

type Despesa struct {
	gorm.Model
	Descricao string `json:"descricao"`
	Valor     string `json:"valor"`
	Data      string `json:"data"`
}
