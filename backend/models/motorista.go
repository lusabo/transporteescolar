package models

import (
	"gorm.io/gorm"
)

type Motorista struct {
	gorm.Model
	Nome    string
	Email   string
	Senha   string
	Escolas []*Escola `gorm:"many2many:motorista_escolas;"`
	Bairros []*Bairro `gorm:"many2many:motorista_bairros;"`
}
